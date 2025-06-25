package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"fmt"
	"strconv"
	"time"
)

type PaymentService interface {
	Create(payment *models.Payment) error
	GetByID(id int) (*models.Payment, error)
	GetByAppointmentID(appointmentID int) (*models.Payment, error)
	GetUserPayments(userID int, limit, offset int) ([]*models.Payment, error)
	List(limit, offset int) ([]*models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id int) error
	ProcessPayment(appointmentID int, paymentMethod models.PaymentMethod, deviceID *int) (*models.Payment, error)
	RefundPayment(paymentID int, reason string) error
	GetTotalRevenue(startDate, endDate time.Time) (float64, error)
	GetPaymentsByStatus(status models.PaymentStatus, limit, offset int) ([]*models.Payment, error)
}

type paymentService struct {
	paymentRepo     repository.PaymentRepository
	appointmentRepo repository.AppointmentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository, appointmentRepo repository.AppointmentRepository) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		appointmentRepo: appointmentRepo,
	}
}

func (s *paymentService) Create(payment *models.Payment) error {
	return s.paymentRepo.Create(payment)
}

func (s *paymentService) GetByID(id int) (*models.Payment, error) {
	payment, err := s.paymentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, fmt.Errorf("payment not found")
	}
	return payment, nil
}

func (s *paymentService) GetByAppointmentID(appointmentID int) (*models.Payment, error) {
	return s.paymentRepo.GetByAppointmentID(appointmentID)
}

func (s *paymentService) GetUserPayments(userID int, limit, offset int) ([]*models.Payment, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.paymentRepo.GetByUserID(userID, limit, offset)
}

func (s *paymentService) List(limit, offset int) ([]*models.Payment, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.paymentRepo.List(limit, offset)
}

func (s *paymentService) Update(payment *models.Payment) error {
	// Verify payment exists
	existing, err := s.paymentRepo.GetByID(payment.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("payment not found")
	}

	return s.paymentRepo.Update(payment)
}

func (s *paymentService) Delete(id int) error {
	// Verify payment exists
	existing, err := s.paymentRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("payment not found")
	}

	// Only allow deletion of failed payments
	if existing.Status == models.PaymentCompleted {
		return fmt.Errorf("cannot delete completed payment")
	}

	return s.paymentRepo.Delete(id)
}

func (s *paymentService) ProcessPayment(appointmentID int, paymentMethod models.PaymentMethod, deviceID *int) (*models.Payment, error) {
	// Get appointment to verify and get amount
	appointment, err := s.appointmentRepo.GetByID(appointmentID)
	if err != nil {
		return nil, fmt.Errorf("appointment not found")
	}

	// Check if payment already exists
	existingPayment, err := s.paymentRepo.GetByAppointmentID(appointmentID)
	if err != nil {
		return nil, err
	}
	if existingPayment != nil && existingPayment.Status == models.PaymentCompleted {
		return nil, fmt.Errorf("appointment already paid")
	}

	// Generate transaction ID
	transactionID := "demo_" + strconv.Itoa(appointmentID) + "_" + strconv.FormatInt(time.Now().Unix(), 10)

	// Create payment
	payment := &models.Payment{
		AppointmentID: appointmentID,
		DeviceID:      deviceID,
		Amount:        appointment.TotalAmount,
		PaymentMethod: paymentMethod,
		TransactionID: transactionID,
		Status:        models.PaymentCompleted, // Demo: always successful
	}

	// Save payment
	err = s.paymentRepo.Create(payment)
	if err != nil {
		return nil, err
	}

	// Update appointment payment status
	err = s.appointmentRepo.UpdatePaymentStatus(appointmentID, models.PaymentCompleted)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) RefundPayment(paymentID int, reason string) error {
	// Get payment
	payment, err := s.paymentRepo.GetByID(paymentID)
	if err != nil {
		return err
	}
	if payment == nil {
		return fmt.Errorf("payment not found")
	}

	// Check if payment can be refunded
	if payment.Status != models.PaymentCompleted {
		return fmt.Errorf("only completed payments can be refunded")
	}

	// Update payment status to refunded
	payment.Status = models.PaymentRefunded

	err = s.paymentRepo.Update(payment)
	if err != nil {
		return err
	}

	// Update appointment payment status
	return s.appointmentRepo.UpdatePaymentStatus(payment.AppointmentID, models.PaymentRefunded)
}

func (s *paymentService) GetTotalRevenue(startDate, endDate time.Time) (float64, error) {
	return s.paymentRepo.GetTotalByDateRange(startDate, endDate)
}

func (s *paymentService) GetPaymentsByStatus(status models.PaymentStatus, limit, offset int) ([]*models.Payment, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.paymentRepo.GetByStatus(status, limit, offset)
}
