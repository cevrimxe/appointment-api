package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
	"time"
)

type DeviceService interface {
	Create(device *models.Device) error
	GetByID(id int) (*models.Device, error)
	Update(device *models.Device) error
	Delete(id int) error
	List() ([]*models.Device, error)
	ListActive() ([]*models.Device, error)
}

type deviceService struct {
	deviceRepo repository.DeviceRepository
}

func NewDeviceService(deviceRepo repository.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepo: deviceRepo,
	}
}

func (s *deviceService) Create(device *models.Device) error {
	if device.Brand == "" {
		return errors.New("device brand is required")
	}

	if device.Name == "" {
		return errors.New("device name is required")
	}

	if device.Price < 0 {
		return errors.New("device price must be positive")
	}

	// Check if device date is not in the future
	if device.DeviceDate.After(time.Now()) {
		return errors.New("device date cannot be in the future")
	}

	// Set default active status
	device.Active = true

	return s.deviceRepo.Create(device)
}

func (s *deviceService) GetByID(id int) (*models.Device, error) {
	if id <= 0 {
		return nil, errors.New("invalid device ID")
	}

	return s.deviceRepo.GetByID(id)
}

func (s *deviceService) Update(device *models.Device) error {
	if device.ID <= 0 {
		return errors.New("invalid device ID")
	}

	if device.Brand == "" {
		return errors.New("device brand is required")
	}

	if device.Name == "" {
		return errors.New("device name is required")
	}

	if device.Price < 0 {
		return errors.New("device price must be positive")
	}

	// Check if device date is not in the future
	if device.DeviceDate.After(time.Now()) {
		return errors.New("device date cannot be in the future")
	}

	// Check if device exists
	existing, err := s.deviceRepo.GetByID(device.ID)
	if err != nil {
		return errors.New("device not found")
	}

	// Update fields
	existing.Brand = device.Brand
	existing.Name = device.Name
	existing.DeviceDate = device.DeviceDate
	existing.Price = device.Price
	existing.Active = device.Active

	return s.deviceRepo.Update(existing)
}

func (s *deviceService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid device ID")
	}

	// Check if device exists
	_, err := s.deviceRepo.GetByID(id)
	if err != nil {
		return errors.New("device not found")
	}

	return s.deviceRepo.Delete(id)
}

func (s *deviceService) List() ([]*models.Device, error) {
	return s.deviceRepo.List()
}

func (s *deviceService) ListActive() ([]*models.Device, error) {
	return s.deviceRepo.ListActive()
}
