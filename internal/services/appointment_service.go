package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
	"time"
)

type AppointmentService interface {
	Create(appointment *models.Appointment) error
	GetByID(id int) (*models.Appointment, error)
	Update(appointment *models.Appointment) error
	Delete(id int) error
	Cancel(id int, userID int) error
	GetByUserID(userID int) ([]*models.Appointment, error)
	List(limit, offset int) ([]*models.Appointment, int, error)
	UpdateStatus(id int, status models.AppointmentStatus) error
	CreateFromRequest(req *models.CreateAppointmentRequest, userID int) (*models.Appointment, error)
	UpdatePaymentStatus(appointmentID int, status models.PaymentStatus) error
	GetTodayCount() (int, error)
	GetMonthlyCount() (int, error)
	GetPreviousTodayCount() (int, error)
	GetPreviousMonthlyCount() (int, error)
}

type appointmentService struct {
	appointmentRepo repository.AppointmentRepository
	serviceRepo     repository.ServiceRepository
	specialistRepo  repository.SpecialistRepository
}

func NewAppointmentService(appointmentRepo repository.AppointmentRepository, serviceRepo repository.ServiceRepository, specialistRepo repository.SpecialistRepository) AppointmentService {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
		serviceRepo:     serviceRepo,
		specialistRepo:  specialistRepo,
	}
}

func (s *appointmentService) CreateFromRequest(req *models.CreateAppointmentRequest, userID int) (*models.Appointment, error) {
	// Validate specialist
	specialist, err := s.specialistRepo.GetByID(req.SpecialistID)
	if err != nil {
		return nil, errors.New("specialist not found")
	}
	if !specialist.Active {
		return nil, errors.New("specialist is not active")
	}

	// Validate service
	service, err := s.serviceRepo.GetByID(req.ServiceID)
	if err != nil {
		return nil, errors.New("service not found")
	}
	if !service.Active {
		return nil, errors.New("service is not active")
	}

	// Check for time conflicts
	hasConflict, err := s.appointmentRepo.CheckConflict(req.SpecialistID, req.AppointmentDate, req.AppointmentTime, nil)
	if err != nil {
		return nil, err
	}
	if hasConflict {
		return nil, errors.New("appointment time is already booked")
	}

	// Check if appointment is in the past
	appointmentDateTime := time.Date(
		req.AppointmentDate.Year(),
		req.AppointmentDate.Month(),
		req.AppointmentDate.Day(),
		req.AppointmentTime.Hour(),
		req.AppointmentTime.Minute(),
		0, 0,
		time.Local,
	)
	if appointmentDateTime.Before(time.Now()) {
		return nil, errors.New("appointment cannot be in the past")
	}

	// Create appointment
	appointment := &models.Appointment{
		UserID:          userID,
		SpecialistID:    req.SpecialistID,
		ServiceID:       req.ServiceID,
		AppointmentDate: req.AppointmentDate,
		AppointmentTime: req.AppointmentTime,
		Status:          models.StatusPending,
		PaymentStatus:   models.PaymentPending,
		TotalAmount:     service.Price,
		Notes:           req.Notes,
	}

	err = s.appointmentRepo.Create(appointment)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *appointmentService) Create(appointment *models.Appointment) error {
	// Validate required fields
	if appointment.UserID <= 0 {
		return errors.New("invalid user ID")
	}
	if appointment.SpecialistID <= 0 {
		return errors.New("invalid specialist ID")
	}
	if appointment.ServiceID <= 0 {
		return errors.New("invalid service ID")
	}

	// Check for time conflicts
	hasConflict, err := s.appointmentRepo.CheckConflict(appointment.SpecialistID, appointment.AppointmentDate, appointment.AppointmentTime, nil)
	if err != nil {
		return err
	}
	if hasConflict {
		return errors.New("appointment time is already booked")
	}

	// Set default status if not provided
	if appointment.Status == "" {
		appointment.Status = models.StatusPending
	}
	if appointment.PaymentStatus == "" {
		appointment.PaymentStatus = models.PaymentPending
	}

	return s.appointmentRepo.Create(appointment)
}

func (s *appointmentService) GetByID(id int) (*models.Appointment, error) {
	if id <= 0 {
		return nil, errors.New("invalid appointment ID")
	}

	return s.appointmentRepo.GetByID(id)
}

func (s *appointmentService) Update(appointment *models.Appointment) error {
	if appointment.ID <= 0 {
		return errors.New("invalid appointment ID")
	}

	// Check if appointment exists
	existing, err := s.appointmentRepo.GetByID(appointment.ID)
	if err != nil {
		return errors.New("appointment not found")
	}

	// Don't allow updating cancelled appointments
	if existing.Status == models.StatusCancelled {
		return errors.New("cannot update cancelled appointment")
	}

	// Check for time conflicts if time/date/specialist changed
	if existing.SpecialistID != appointment.SpecialistID ||
		!existing.AppointmentDate.Equal(appointment.AppointmentDate) ||
		!existing.AppointmentTime.Equal(appointment.AppointmentTime) {

		hasConflict, err := s.appointmentRepo.CheckConflict(appointment.SpecialistID, appointment.AppointmentDate, appointment.AppointmentTime, &appointment.ID)
		if err != nil {
			return err
		}
		if hasConflict {
			return errors.New("appointment time is already booked")
		}
	}

	return s.appointmentRepo.Update(appointment)
}

func (s *appointmentService) Cancel(id int, userID int) error {
	if id <= 0 {
		return errors.New("invalid appointment ID")
	}

	// Check if appointment exists and belongs to user
	appointment, err := s.appointmentRepo.GetByID(id)
	if err != nil {
		return errors.New("appointment not found")
	}

	if appointment.UserID != userID {
		return errors.New("unauthorized to cancel this appointment")
	}

	if appointment.Status == models.StatusCancelled {
		return errors.New("appointment is already cancelled")
	}

	if appointment.Status == models.StatusCompleted {
		return errors.New("cannot cancel completed appointment")
	}

	return s.appointmentRepo.UpdateStatus(id, models.StatusCancelled)
}

func (s *appointmentService) GetByUserID(userID int) ([]*models.Appointment, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return s.appointmentRepo.GetByUserID(userID)
}

func (s *appointmentService) List(limit, offset int) ([]*models.Appointment, int, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return s.appointmentRepo.List(limit, offset)
}

func (s *appointmentService) UpdateStatus(id int, status models.AppointmentStatus) error {
	if id <= 0 {
		return errors.New("invalid appointment ID")
	}

	// Check if appointment exists
	_, err := s.appointmentRepo.GetByID(id)
	if err != nil {
		return errors.New("appointment not found")
	}

	// Validate status
	validStatuses := map[models.AppointmentStatus]bool{
		models.StatusPending:   true,
		models.StatusConfirmed: true,
		models.StatusCompleted: true,
		models.StatusCancelled: true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return s.appointmentRepo.UpdateStatus(id, status)
}

func (s *appointmentService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid appointment ID")
	}

	// Check if appointment exists
	_, err := s.appointmentRepo.GetByID(id)
	if err != nil {
		return errors.New("appointment not found")
	}

	return s.appointmentRepo.Delete(id)
}

func (s *appointmentService) UpdatePaymentStatus(appointmentID int, status models.PaymentStatus) error {
	return s.appointmentRepo.UpdatePaymentStatus(appointmentID, status)
}

func (s *appointmentService) GetTodayCount() (int, error) {
	// Use List method with large limit to get all appointments and filter by today
	appointments, _, err := s.appointmentRepo.List(10000, 0)
	if err != nil {
		return 0, err
	}

	today := time.Now().Format("2006-01-02")
	count := 0
	for _, appointment := range appointments {
		appointmentDate := appointment.AppointmentDate.Format("2006-01-02")
		if appointmentDate == today {
			count++
		}
	}
	return count, nil
}

func (s *appointmentService) GetMonthlyCount() (int, error) {
	// Use List method to get all appointments and filter by this month
	appointments, _, err := s.appointmentRepo.List(10000, 0)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()
	count := 0
	for _, appointment := range appointments {
		if appointment.AppointmentDate.Month() == currentMonth && appointment.AppointmentDate.Year() == currentYear {
			count++
		}
	}
	return count, nil
}

func (s *appointmentService) GetPreviousTodayCount() (int, error) {
	// Use List method to get all appointments and filter by yesterday
	appointments, _, err := s.appointmentRepo.List(10000, 0)
	if err != nil {
		return 0, err
	}

	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	count := 0
	for _, appointment := range appointments {
		appointmentDate := appointment.AppointmentDate.Format("2006-01-02")
		if appointmentDate == yesterday {
			count++
		}
	}
	return count, nil
}

func (s *appointmentService) GetPreviousMonthlyCount() (int, error) {
	// Use List method to get all appointments and filter by previous month
	appointments, _, err := s.appointmentRepo.List(10000, 0)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	prevMonth := now.AddDate(0, -1, 0)
	count := 0
	for _, appointment := range appointments {
		if appointment.AppointmentDate.Month() == prevMonth.Month() && appointment.AppointmentDate.Year() == prevMonth.Year() {
			count++
		}
	}
	return count, nil
}
