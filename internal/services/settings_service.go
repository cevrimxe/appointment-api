package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
	"strconv"
)

type SettingsService interface {
	GetByKey(key string) (*models.Setting, error)
	Update(setting *models.Setting) error
	List() ([]*models.Setting, error)
	UpdateAppointmentDuration(minutes int) error
}

type settingsService struct {
	settingsRepo repository.SettingsRepository
	serviceRepo  repository.ServiceRepository
}

func NewSettingsService(settingsRepo repository.SettingsRepository, serviceRepo repository.ServiceRepository) SettingsService {
	return &settingsService{
		settingsRepo: settingsRepo,
		serviceRepo:  serviceRepo,
	}
}

func (s *settingsService) GetByKey(key string) (*models.Setting, error) {
	if key == "" {
		return nil, errors.New("setting key is required")
	}

	return s.settingsRepo.GetByKey(key)
}

func (s *settingsService) Update(setting *models.Setting) error {
	if setting.Key == "" {
		return errors.New("setting key is required")
	}

	if setting.Value == "" {
		return errors.New("setting value is required")
	}

	// Check if setting exists
	_, err := s.settingsRepo.GetByKey(setting.Key)
	if err != nil {
		return errors.New("setting not found")
	}

	return s.settingsRepo.UpdateByKey(setting.Key, setting.Value, setting.Description)
}

func (s *settingsService) List() ([]*models.Setting, error) {
	return s.settingsRepo.List()
}

func (s *settingsService) UpdateAppointmentDuration(minutes int) error {
	if minutes <= 0 {
		return errors.New("appointment duration must be positive")
	}

	if minutes > 480 { // Max 8 hours
		return errors.New("appointment duration cannot exceed 480 minutes")
	}

	// Update the setting using UpdateByKey method
	return s.settingsRepo.UpdateByKey("appointment_duration", strconv.Itoa(minutes), "Default appointment duration in minutes")
}
