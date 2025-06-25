package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SpecialistService interface {
	Create(specialist *models.Specialist) error
	GetByID(id int) (*models.Specialist, error)
	Update(specialist *models.Specialist) error
	Delete(id int) error
	List() ([]*models.Specialist, error)
	ListActive() ([]*models.Specialist, error)
	GetWorkingHours(specialistID int) ([]*models.WorkingHour, error)
	UpdateWorkingHours(specialistID int, workingHours []*models.WorkingHour) error
	GetAvailableSlots(specialistID int, date string) ([]string, error)
}

type specialistService struct {
	specialistRepo  repository.SpecialistRepository
	appointmentRepo repository.AppointmentRepository
	settingsRepo    repository.SettingsRepository
}

func NewSpecialistService(specialistRepo repository.SpecialistRepository, appointmentRepo repository.AppointmentRepository, settingsRepo repository.SettingsRepository) SpecialistService {
	return &specialistService{
		specialistRepo:  specialistRepo,
		appointmentRepo: appointmentRepo,
		settingsRepo:    settingsRepo,
	}
}

func (s *specialistService) Create(specialist *models.Specialist) error {
	if specialist.Name == "" {
		return errors.New("specialist name is required")
	}

	if specialist.Email == "" {
		return errors.New("specialist email is required")
	}

	// Set default active status
	specialist.Active = true

	return s.specialistRepo.Create(specialist)
}

func (s *specialistService) GetByID(id int) (*models.Specialist, error) {
	if id <= 0 {
		return nil, errors.New("invalid specialist ID")
	}

	return s.specialistRepo.GetByID(id)
}

func (s *specialistService) Update(specialist *models.Specialist) error {
	if specialist.ID <= 0 {
		return errors.New("invalid specialist ID")
	}

	if specialist.Name == "" {
		return errors.New("specialist name is required")
	}

	if specialist.Email == "" {
		return errors.New("specialist email is required")
	}

	// Check if specialist exists
	existing, err := s.specialistRepo.GetByID(specialist.ID)
	if err != nil {
		return errors.New("specialist not found")
	}

	// Update fields
	existing.Name = specialist.Name
	existing.Email = specialist.Email
	existing.Phone = specialist.Phone
	existing.Active = specialist.Active

	return s.specialistRepo.Update(existing)
}

func (s *specialistService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid specialist ID")
	}

	// Check if specialist exists
	_, err := s.specialistRepo.GetByID(id)
	if err != nil {
		return errors.New("specialist not found")
	}

	return s.specialistRepo.Delete(id)
}

func (s *specialistService) List() ([]*models.Specialist, error) {
	return s.specialistRepo.List()
}

func (s *specialistService) ListActive() ([]*models.Specialist, error) {
	return s.specialistRepo.ListActive()
}

func (s *specialistService) GetWorkingHours(specialistID int) ([]*models.WorkingHour, error) {
	if specialistID <= 0 {
		return nil, errors.New("invalid specialist ID")
	}

	// Check if specialist exists
	_, err := s.specialistRepo.GetByID(specialistID)
	if err != nil {
		return nil, errors.New("specialist not found")
	}

	return s.specialistRepo.GetWorkingHours(specialistID)
}

func (s *specialistService) UpdateWorkingHours(specialistID int, workingHours []*models.WorkingHour) error {
	if specialistID <= 0 {
		return errors.New("invalid specialist ID")
	}

	// Check if specialist exists
	_, err := s.specialistRepo.GetByID(specialistID)
	if err != nil {
		return errors.New("specialist not found")
	}

	// Validate working hours
	for i, wh := range workingHours {
		if wh.DayOfWeek < 0 || wh.DayOfWeek > 6 {
			return fmt.Errorf("invalid day of week %d for working hour %d", wh.DayOfWeek, i+1)
		}

		if !isValidTimeFormat(wh.StartTime) {
			return fmt.Errorf("invalid start time format for working hour %d", i+1)
		}

		if !isValidTimeFormat(wh.EndTime) {
			return fmt.Errorf("invalid end time format for working hour %d", i+1)
		}

		// Validate that start time is before end time
		if !isTimeBeforeTime(wh.StartTime, wh.EndTime) {
			return fmt.Errorf("start time must be before end time for working hour %d", i+1)
		}

		// Set specialist ID
		wh.SpecialistID = specialistID
	}

	return s.specialistRepo.UpdateWorkingHours(specialistID, workingHours)
}

// Helper function to validate time format (HH:MM)
func isValidTimeFormat(timeStr string) bool {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return false
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil || hour < 0 || hour > 23 {
		return false
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil || minute < 0 || minute > 59 {
		return false
	}

	return true
}

// Helper function to check if time1 is before time2
func isTimeBeforeTime(time1, time2 string) bool {
	parts1 := strings.Split(time1, ":")
	parts2 := strings.Split(time2, ":")

	hour1, _ := strconv.Atoi(parts1[0])
	minute1, _ := strconv.Atoi(parts1[1])
	hour2, _ := strconv.Atoi(parts2[0])
	minute2, _ := strconv.Atoi(parts2[1])

	total1 := hour1*60 + minute1
	total2 := hour2*60 + minute2

	return total1 < total2
}

func (s *specialistService) GetAvailableSlots(specialistID int, date string) ([]string, error) {
	if specialistID <= 0 {
		return nil, errors.New("invalid specialist ID")
	}

	// Check if specialist exists
	_, err := s.specialistRepo.GetByID(specialistID)
	if err != nil {
		return nil, errors.New("specialist not found")
	}

	// Parse the date
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	// Get day of week (0=Sunday, 1=Monday, etc.)
	dayOfWeek := int(parsedDate.Weekday())

	// Get working hours for this specialist
	workingHours, err := s.specialistRepo.GetWorkingHours(specialistID)
	if err != nil {
		return nil, errors.New("failed to get working hours")
	}

	// Find working hours for the requested day
	var dayWorkingHours *models.WorkingHour
	for _, wh := range workingHours {
		if wh.DayOfWeek == dayOfWeek && wh.Active {
			dayWorkingHours = wh
			break
		}
	}

	// If no working hours for this day, return empty slots
	if dayWorkingHours == nil {
		return []string{}, nil
	}

	// Get appointment duration from settings (default 60 minutes)
	var appointmentDuration int = 60
	durationSetting, err := s.settingsRepo.GetByKey("appointment_duration")
	if err == nil && durationSetting.Value != "" {
		if duration, parseErr := strconv.Atoi(durationSetting.Value); parseErr == nil && duration > 0 {
			appointmentDuration = duration
		}
	}

	// Generate time slots based on working hours and appointment duration
	slots := []string{}
	startTime := dayWorkingHours.StartTime
	endTime := dayWorkingHours.EndTime

	// Parse start and end times - handle both HH:MM and timestamp formats
	var startHour, startMinute, endHour, endMinute int

	// Try parsing as timestamp format first (0000-01-01T09:00:00Z)
	if parsedStartTime, parseErr := time.Parse("2006-01-02T15:04:05Z", startTime); parseErr == nil {
		startHour = parsedStartTime.Hour()
		startMinute = parsedStartTime.Minute()
	} else {
		// Fall back to HH:MM format
		startParts := strings.Split(startTime, ":")
		if len(startParts) != 2 {
			return nil, errors.New("invalid start time format")
		}
		startHour, _ = strconv.Atoi(startParts[0])
		startMinute, _ = strconv.Atoi(startParts[1])
	}

	if parsedEndTime, parseErr := time.Parse("2006-01-02T15:04:05Z", endTime); parseErr == nil {
		endHour = parsedEndTime.Hour()
		endMinute = parsedEndTime.Minute()
	} else {
		// Fall back to HH:MM format
		endParts := strings.Split(endTime, ":")
		if len(endParts) != 2 {
			return nil, errors.New("invalid end time format")
		}
		endHour, _ = strconv.Atoi(endParts[0])
		endMinute, _ = strconv.Atoi(endParts[1])
	}

	// Convert to minutes for easier calculation
	startTotalMinutes := startHour*60 + startMinute
	endTotalMinutes := endHour*60 + endMinute

	// Generate slots with appointment duration intervals
	for currentMinutes := startTotalMinutes; currentMinutes+appointmentDuration <= endTotalMinutes; currentMinutes += appointmentDuration {
		hour := currentMinutes / 60
		minute := currentMinutes % 60
		slotTime := fmt.Sprintf("%02d:%02d", hour, minute)
		slots = append(slots, slotTime)
	}

	// Get existing appointments for this specialist on this date
	existingAppointments, err := s.getAppointmentsBySpecialistAndDate(specialistID, parsedDate)
	if err != nil {
		// Log error but continue (return all available slots)
		fmt.Printf("Warning: failed to get existing appointments: %v\n", err)
		return slots, nil
	}

	// Remove booked slots
	availableSlots := []string{}
	for _, slot := range slots {
		isBooked := false
		for _, appointment := range existingAppointments {
			appointmentTime := appointment.AppointmentTime.Format("15:04")
			if appointmentTime == slot && appointment.Status != models.StatusCancelled {
				isBooked = true
				break
			}
		}
		if !isBooked {
			availableSlots = append(availableSlots, slot)
		}
	}

	return availableSlots, nil
}

// Helper method to get appointments by specialist and date
func (s *specialistService) getAppointmentsBySpecialistAndDate(specialistID int, date time.Time) ([]*models.Appointment, error) {
	return s.appointmentRepo.GetBySpecialistID(specialistID, &date)
}
