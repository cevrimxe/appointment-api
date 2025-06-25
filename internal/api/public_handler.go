package api

import (
	"appointment-api/internal/middleware"
	"appointment-api/internal/services"
	"net/http"
	"strconv"
	"time"

	"appointment-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PublicHandler struct {
	categoryService    services.CategoryService
	serviceService     services.ServiceService
	specialistService  services.SpecialistService
	appointmentService services.AppointmentService
	paymentService     services.PaymentService
	contactService     services.ContactService
	validator          *validator.Validate
}

func NewPublicHandler(categoryService services.CategoryService, serviceService services.ServiceService, specialistService services.SpecialistService, appointmentService services.AppointmentService, paymentService services.PaymentService, contactService services.ContactService, validator *validator.Validate) *PublicHandler {
	return &PublicHandler{
		categoryService:    categoryService,
		serviceService:     serviceService,
		specialistService:  specialistService,
		appointmentService: appointmentService,
		paymentService:     paymentService,
		contactService:     contactService,
		validator:          validator,
	}
}

// Categories endpoints
func (h *PublicHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.ListActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    categories,
	})
}

func (h *PublicHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid category ID",
		})
		return
	}

	category, err := h.categoryService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
	})
}

// Services endpoints
func (h *PublicHandler) GetServices(c *gin.Context) {
	services, err := h.serviceService.ListActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch services",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    services,
	})
}

func (h *PublicHandler) GetServiceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid service ID",
		})
		return
	}

	service, err := h.serviceService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Service not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    service,
	})
}

// Get active specialists (public)
func (h *PublicHandler) GetSpecialists(c *gin.Context) {
	specialists, err := h.specialistService.ListActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch specialists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    specialists,
	})
}

func (h *PublicHandler) GetSpecialistByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid specialist ID",
		})
		return
	}

	specialist, err := h.specialistService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Specialist not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    specialist,
	})
}

func (h *PublicHandler) GetServicesByCategory(c *gin.Context) {
	idStr := c.Param("id")
	categoryID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid category ID",
		})
		return
	}

	services, err := h.serviceService.ListByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    services,
	})
}

func (h *PublicHandler) GetSpecialistWorkingHours(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid specialist ID",
		})
		return
	}

	workingHours, err := h.specialistService.GetWorkingHours(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    workingHours,
	})
}

func (h *PublicHandler) GetSpecialistAvailableSlots(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid specialist ID",
		})
		return
	}

	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Date parameter is required (YYYY-MM-DD format)",
		})
		return
	}

	availableSlots, err := h.specialistService.GetAvailableSlots(id, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    availableSlots,
	})
}

func (h *PublicHandler) ContactMessage(c *gin.Context) {
	var req models.ContactMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	message, err := h.contactService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to send message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    message,
		"message": "Your message has been sent successfully",
	})
}

func (h *PublicHandler) CreateAppointment(c *gin.Context) {
	// Get current user from middleware
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	currentUser := user

	var req models.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	appointment, err := h.appointmentService.CreateFromRequest(&req, currentUser.ID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "specialist not found" || err.Error() == "service not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "appointment time is already booked" ||
			err.Error() == "appointment cannot be in the past" ||
			err.Error() == "specialist is not active" ||
			err.Error() == "service is not active" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    appointment,
		"message": "Appointment created successfully",
	})
}

func (h *PublicHandler) GetUserAppointments(c *gin.Context) {
	// Get current user from middleware
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	currentUser := user

	appointments, err := h.appointmentService.GetByUserID(currentUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    appointments,
	})
}

func (h *PublicHandler) GetAppointmentByID(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	appointmentIDStr := c.Param("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid appointment ID",
		})
		return
	}

	appointment, err := h.appointmentService.GetByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Appointment not found",
		})
		return
	}

	// Verify ownership
	if appointment.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied to this appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    appointment,
	})
}

func (h *PublicHandler) UpdateAppointment(c *gin.Context) {
	// Get current user from middleware
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	currentUser := user

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid appointment ID",
		})
		return
	}

	// Get existing appointment
	existing, err := h.appointmentService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Appointment not found",
		})
		return
	}

	// Check if appointment belongs to current user
	if existing.UserID != currentUser.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied",
		})
		return
	}

	var req struct {
		SpecialistID    *int       `json:"specialist_id"`
		ServiceID       *int       `json:"service_id"`
		AppointmentDate *time.Time `json:"appointment_date"`
		AppointmentTime *time.Time `json:"appointment_time"`
		Notes           *string    `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	// Update only provided fields (non-nil pointers)
	if req.SpecialistID != nil {
		existing.SpecialistID = *req.SpecialistID
	}

	if req.ServiceID != nil {
		existing.ServiceID = *req.ServiceID
	}

	if req.AppointmentDate != nil {
		existing.AppointmentDate = *req.AppointmentDate
	}

	if req.AppointmentTime != nil {
		existing.AppointmentTime = *req.AppointmentTime
	}

	if req.Notes != nil {
		existing.Notes = *req.Notes
	}

	err = h.appointmentService.Update(existing)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "appointment not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "appointment time is already booked" ||
			err.Error() == "cannot update cancelled appointment" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existing,
		"message": "Appointment updated successfully",
	})
}

func (h *PublicHandler) CancelAppointment(c *gin.Context) {
	// Get current user from middleware
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	currentUser := user

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid appointment ID",
		})
		return
	}

	err = h.appointmentService.Cancel(id, currentUser.ID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "appointment not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "unauthorized to cancel this appointment" {
			statusCode = http.StatusForbidden
		} else if err.Error() == "appointment is already cancelled" ||
			err.Error() == "cannot cancel completed appointment" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Appointment cancelled successfully",
	})
}

func (h *PublicHandler) PayAppointment(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	appointmentIDStr := c.Param("id")
	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid appointment ID",
		})
		return
	}

	var req struct {
		PaymentMethod string `json:"payment_method" validate:"required"`
		CardToken     string `json:"card_token"`
		DeviceID      *int   `json:"device_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request format",
		})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// Convert payment method
	var paymentMethod models.PaymentMethod
	switch req.PaymentMethod {
	case "credit_card":
		paymentMethod = models.PaymentMethodCreditCard
	case "cash":
		paymentMethod = models.PaymentMethodCash
	case "transfer":
		paymentMethod = models.PaymentMethodTransfer
	default:
		paymentMethod = models.PaymentMethodCreditCard
	}

	// First verify appointment ownership
	appointment, err := h.appointmentService.GetByID(appointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Appointment not found",
		})
		return
	}

	// Verify ownership
	if appointment.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied to this appointment",
		})
		return
	}

	// Process payment using service
	payment, err := h.paymentService.ProcessPayment(appointmentID, paymentMethod, req.DeviceID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "appointment not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "appointment already paid" {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payment,
		"message": "Payment processed successfully",
	})
}

func (h *PublicHandler) GetUserPayments(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	// Parse query parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	payments, err := h.paymentService.GetUserPayments(user.ID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch payments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payments,
		"pagination": gin.H{
			"limit":  limit,
			"offset": offset,
			"count":  len(payments),
		},
	})
}

func (h *PublicHandler) GetPaymentByID(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "User not authenticated",
		})
		return
	}

	paymentIDStr := c.Param("id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid payment ID",
		})
		return
	}

	payment, err := h.paymentService.GetByID(paymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Payment not found",
		})
		return
	}

	// Get appointment to verify ownership
	appointment, err := h.appointmentService.GetByID(payment.AppointmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Payment not found",
		})
		return
	}

	// Verify ownership through appointment
	if appointment.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Access denied to this payment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payment,
	})
}
