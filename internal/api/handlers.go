package api

import (
	"appointment-api/internal/config"
	"appointment-api/internal/middleware"
	"appointment-api/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	Auth   *AuthHandler
	Public *PublicHandler
	Admin  *AdminHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	validate := validator.New()
	return &Handlers{
		Auth:   NewAuthHandler(svc.Auth),
		Public: NewPublicHandler(svc.Category, svc.Service, svc.Specialist, svc.Appointment, svc.Payment, svc.Contact, validate),
		Admin:  NewAdminHandler(svc.Category, svc.Service, svc.Device, svc.Settings, svc.Auth, svc.User, svc.Specialist, svc.Appointment, svc.Payment, svc.Contact),
	}
}

func SetupRoutes(router *gin.Engine, handlers *Handlers, svc *services.Services, mainDB *sql.DB, cfg *config.Config) {
	// Remove request logging to keep logs clean

	// Health check (without tenant middleware)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Service is running",
		})
	})

	// Cache stats endpoint (for monitoring)
	router.GET("/cache/stats", func(c *gin.Context) {
		count, domains := svc.TenantCache.GetCacheStats()
		c.JSON(200, gin.H{
			"tenant_count": count,
			"domains":      domains,
		})
	})

	api := router.Group("/api")
	api.Use(middleware.TenantMiddleware(svc.Tenant, svc.TenantCache, mainDB))
	// Simple tenant context middleware
	api.Use(func(c *gin.Context) {
		c.Next()
	})
	{

		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Auth.Register)
			auth.POST("/login", handlers.Auth.Login)
			auth.POST("/admin/login", handlers.Auth.AdminLogin)
			auth.POST("/forgot-password", handlers.Auth.ForgotPassword)
			auth.POST("/reset-password", handlers.Auth.ResetPassword)
		}

		// Categories routes (public)
		categories := api.Group("/categories")
		{
			categories.GET("", handlers.Public.GetCategories)
			categories.GET("/:id", handlers.Public.GetCategoryByID)
			categories.GET("/:id/services", handlers.Public.GetServicesByCategory)
		}

		// Services routes (public)
		services := api.Group("/services")
		{
			services.GET("", handlers.Public.GetServices)
			services.GET("/:id", handlers.Public.GetServiceByID)
		}

		// Specialists routes (public)
		specialists := api.Group("/specialists")
		{
			specialists.GET("", handlers.Public.GetSpecialists)
			specialists.GET("/:id", handlers.Public.GetSpecialistByID)
			specialists.GET("/:id/working-hours", handlers.Public.GetSpecialistWorkingHours)
			specialists.GET("/:id/available-slots", handlers.Public.GetSpecialistAvailableSlots)
		}

		// Contact route (public)
		api.POST("/contact", handlers.Public.ContactMessage)

		// Public routes (categories & services)
		public := api.Group("/public")
		{
			publicCategories := public.Group("/categories")
			{
				publicCategories.GET("", handlers.Public.GetCategories)
				publicCategories.GET("/:id", handlers.Public.GetCategoryByID)
				publicCategories.GET("/:id/services", handlers.Public.GetServicesByCategory)
			}

			publicServices := public.Group("/services")
			{
				publicServices.GET("", handlers.Public.GetServices)
				publicServices.GET("/:id", handlers.Public.GetServiceByID)
			}

			publicSpecialists := public.Group("/specialists")
			{
				publicSpecialists.GET("", handlers.Public.GetSpecialists)
				publicSpecialists.GET("/:id", handlers.Public.GetSpecialistByID)
			}
		}

		// User routes (authenticated)
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware(svc.Auth))
		{
			user.GET("/profile", handlers.Auth.GetProfile)
			user.PUT("/profile", handlers.Auth.UpdateProfile)
			user.PUT("/change-password", handlers.Auth.ChangePassword)
		}

		// Appointments routes (authenticated)
		appointments := api.Group("/appointments")
		appointments.Use(middleware.AuthMiddleware(svc.Auth))
		{
			appointments.POST("", handlers.Public.CreateAppointment)
			appointments.GET("", handlers.Public.GetUserAppointments)
			appointments.GET("/:id", handlers.Public.GetAppointmentByID)
			appointments.PUT("/:id", handlers.Public.UpdateAppointment)
			appointments.DELETE("/:id", handlers.Public.CancelAppointment)
			appointments.POST("/:id/payment", handlers.Public.PayAppointment)
		}

		// Payments routes (authenticated)
		payments := api.Group("/payments")
		payments.Use(middleware.AuthMiddleware(svc.Auth))
		{
			payments.GET("", handlers.Public.GetUserPayments)
			payments.GET("/:id", handlers.Public.GetPaymentByID)
		}

		// Admin routes (admin only)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(svc.Auth))
		admin.Use(middleware.AdminMiddleware())
		{
			// Dashboard Stats
			admin.GET("/stats", handlers.Admin.GetStats)
			admin.GET("/dashboard/stats", handlers.Admin.GetDashboardStats)

			// Categories CRUD
			adminCategories := admin.Group("/categories")
			{
				adminCategories.POST("", handlers.Admin.CreateCategory)
				adminCategories.GET("", handlers.Admin.GetCategories)
				adminCategories.PUT("/:id", handlers.Admin.UpdateCategory)
				adminCategories.DELETE("/:id", handlers.Admin.DeleteCategory)
			}

			// Services CRUD
			adminServices := admin.Group("/services")
			{
				adminServices.POST("", handlers.Admin.CreateService)
				adminServices.GET("", handlers.Admin.GetServices)
				adminServices.PUT("/:id", handlers.Admin.UpdateService)
				adminServices.DELETE("/:id", handlers.Admin.DeleteService)
			}

			// Devices CRUD
			adminDevices := admin.Group("/devices")
			{
				adminDevices.POST("", handlers.Admin.CreateDevice)
				adminDevices.GET("", handlers.Admin.GetDevices)
				adminDevices.PUT("/:id", handlers.Admin.UpdateDevice)
				adminDevices.DELETE("/:id", handlers.Admin.DeleteDevice)
			}

			// Settings Management
			adminSettings := admin.Group("/settings")
			{
				adminSettings.GET("", handlers.Admin.GetSettings)
				adminSettings.PUT("/:key", handlers.Admin.UpdateSetting)
				adminSettings.PUT("/appointment-duration", handlers.Admin.UpdateAppointmentDuration)
			}

			// Users Management
			adminUsers := admin.Group("/users")
			{
				adminUsers.GET("", handlers.Admin.GetUsers)
				adminUsers.POST("", handlers.Admin.CreateUser)
				adminUsers.PUT("/:id", handlers.Admin.UpdateUser)
				adminUsers.DELETE("/:id", handlers.Admin.DeleteUser)
				adminUsers.PUT("/:id/role", handlers.Admin.UpdateUserRole)
			}

			// Specialists Management
			adminSpecialists := admin.Group("/specialists")
			{
				adminSpecialists.GET("", handlers.Admin.GetSpecialists)
				adminSpecialists.POST("", handlers.Admin.CreateSpecialist)
				adminSpecialists.PUT("/:id", handlers.Admin.UpdateSpecialist)
				adminSpecialists.DELETE("/:id", handlers.Admin.DeleteSpecialist)
				adminSpecialists.GET("/:id/working-hours", handlers.Admin.GetSpecialistWorkingHours)
				adminSpecialists.PUT("/:id/working-hours", handlers.Admin.UpdateSpecialistWorkingHours)
			}

			// Appointments Management
			adminAppointments := admin.Group("/appointments")
			{
				adminAppointments.GET("", handlers.Admin.GetAppointments)
				adminAppointments.POST("", handlers.Admin.CreateAppointment)
				adminAppointments.PUT("/:id", handlers.Admin.UpdateAppointment)
				adminAppointments.DELETE("/:id", handlers.Admin.DeleteAppointment)
				adminAppointments.PUT("/:id/status", handlers.Admin.UpdateAppointmentStatus)
			}

			// Payments Management
			adminPayments := admin.Group("/payments")
			{
				adminPayments.GET("", handlers.Admin.GetPayments)
				adminPayments.POST("", handlers.Admin.CreatePayment)
				adminPayments.PUT("/:id", handlers.Admin.UpdatePayment)
				adminPayments.DELETE("/:id", handlers.Admin.DeletePayment)
			}

			// Contact Messages Management
			adminContactMessages := admin.Group("/contact-messages")
			{
				adminContactMessages.GET("", handlers.Admin.GetContactMessages)
				adminContactMessages.PUT("/:id/read", handlers.Admin.MarkContactMessageRead)
				adminContactMessages.DELETE("/:id", handlers.Admin.DeleteContactMessage)
			}

			// Reports
			adminReports := admin.Group("/reports")
			{
				adminReports.GET("/sales", handlers.Admin.GetSalesReports)
				adminReports.GET("/payments", handlers.Admin.GetPaymentReports)
				adminReports.GET("/appointments", handlers.Admin.GetAppointmentReports)
			}
		}
	}
}
