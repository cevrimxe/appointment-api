package api

import (
	"appointment-api/internal/models"
	"appointment-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	categoryService    services.CategoryService
	serviceService     services.ServiceService
	deviceService      services.DeviceService
	settingsService    services.SettingsService
	authService        services.AuthService
	userService        services.UserService
	specialistService  services.SpecialistService
	appointmentService services.AppointmentService
	paymentService     services.PaymentService
	contactService     services.ContactService
}

func NewAdminHandler(
	categoryService services.CategoryService,
	serviceService services.ServiceService,
	deviceService services.DeviceService,
	settingsService services.SettingsService,
	authService services.AuthService,
	userService services.UserService,
	specialistService services.SpecialistService,
	appointmentService services.AppointmentService,
	paymentService services.PaymentService,
	contactService services.ContactService,
) *AdminHandler {
	return &AdminHandler{
		categoryService:    categoryService,
		serviceService:     serviceService,
		deviceService:      deviceService,
		settingsService:    settingsService,
		authService:        authService,
		userService:        userService,
		specialistService:  specialistService,
		appointmentService: appointmentService,
		paymentService:     paymentService,
		contactService:     contactService,
	}
}

// Categories
func (h *AdminHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	err := h.categoryService.Create(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    category,
		"message": "Category created successfully",
	})
}

func (h *AdminHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    categories,
	})
}

func (h *AdminHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid category ID",
		})
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	category.ID = id
	err = h.categoryService.Update(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
		"message": "Category updated successfully",
	})
}

func (h *AdminHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid category ID",
		})
		return
	}

	err = h.categoryService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category deleted successfully",
	})
}

// Services
func (h *AdminHandler) CreateService(c *gin.Context) {
	// TODO: Implement service creation
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Service creation - to be implemented",
	})
}

func (h *AdminHandler) GetServices(c *gin.Context) {
	// TODO: Implement service listing
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Service listing - to be implemented",
	})
}

func (h *AdminHandler) UpdateService(c *gin.Context) {
	// TODO: Implement service update
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Service update - to be implemented",
	})
}

func (h *AdminHandler) DeleteService(c *gin.Context) {
	// TODO: Implement service deletion
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Service deletion - to be implemented",
	})
}

// Placeholder methods for other admin endpoints
func (h *AdminHandler) CreateDevice(c *gin.Context)  { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetDevices(c *gin.Context)    { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateDevice(c *gin.Context)  { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeleteDevice(c *gin.Context)  { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetSettings(c *gin.Context)   { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateSetting(c *gin.Context) { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateAppointmentDuration(c *gin.Context) {
	c.JSON(200, gin.H{"message": "TODO"})
}
func (h *AdminHandler) GetUsers(c *gin.Context)         { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) CreateUser(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateUser(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeleteUser(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateUserRole(c *gin.Context)   { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetSpecialists(c *gin.Context)   { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) CreateSpecialist(c *gin.Context) { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateSpecialist(c *gin.Context) { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeleteSpecialist(c *gin.Context) { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetSpecialistWorkingHours(c *gin.Context) {
	c.JSON(200, gin.H{"message": "TODO"})
}
func (h *AdminHandler) UpdateSpecialistWorkingHours(c *gin.Context) {
	c.JSON(200, gin.H{"message": "TODO"})
}
func (h *AdminHandler) GetAppointments(c *gin.Context)         { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) CreateAppointment(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateAppointment(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeleteAppointment(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdateAppointmentStatus(c *gin.Context) { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetPayments(c *gin.Context)             { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) CreatePayment(c *gin.Context)           { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) UpdatePayment(c *gin.Context)           { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeletePayment(c *gin.Context)           { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetContactMessages(c *gin.Context)      { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) MarkContactMessageRead(c *gin.Context)  { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) DeleteContactMessage(c *gin.Context)    { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetSalesReports(c *gin.Context)         { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetPaymentReports(c *gin.Context)       { c.JSON(200, gin.H{"message": "TODO"}) }
func (h *AdminHandler) GetAppointmentReports(c *gin.Context)   { c.JSON(200, gin.H{"message": "TODO"}) }

// Dashboard Stats
func (h *AdminHandler) GetStats(c *gin.Context) {
	// Demo stats - gerçek veriler için servis methodları gerekli
	stats := gin.H{
		"today_appointments":   12,
		"total_revenue":        45000.50,
		"total_users":          156,
		"pending_appointments": 8,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}
