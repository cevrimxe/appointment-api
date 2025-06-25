package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
)

type ServiceService interface {
	Create(service *models.Service) error
	GetByID(id int) (*models.Service, error)
	Update(service *models.Service) error
	Delete(id int) error
	List() ([]*models.Service, error)
	ListActive() ([]*models.Service, error)
	ListByCategory(categoryID int) ([]*models.Service, error)
}

type serviceService struct {
	serviceRepo  repository.ServiceRepository
	categoryRepo repository.CategoryRepository
}

func NewServiceService(serviceRepo repository.ServiceRepository, categoryRepo repository.CategoryRepository) ServiceService {
	return &serviceService{
		serviceRepo:  serviceRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *serviceService) Create(service *models.Service) error {
	if service.Name == "" {
		return errors.New("service name is required")
	}

	if service.Price < 0 {
		return errors.New("service price must be positive")
	}

	// Validate category if provided
	if service.CategoryID != nil {
		_, err := s.categoryRepo.GetByID(*service.CategoryID)
		if err != nil {
			return errors.New("invalid category")
		}
	}

	// Set default active status
	service.Active = true

	return s.serviceRepo.Create(service)
}

func (s *serviceService) GetByID(id int) (*models.Service, error) {
	if id <= 0 {
		return nil, errors.New("invalid service ID")
	}

	return s.serviceRepo.GetByID(id)
}

func (s *serviceService) Update(service *models.Service) error {
	if service.ID <= 0 {
		return errors.New("invalid service ID")
	}

	if service.Name == "" {
		return errors.New("service name is required")
	}

	if service.Price < 0 {
		return errors.New("service price must be positive")
	}

	// Check if service exists
	existing, err := s.serviceRepo.GetByID(service.ID)
	if err != nil {
		return errors.New("service not found")
	}

	// Validate category if provided
	if service.CategoryID != nil {
		_, err := s.categoryRepo.GetByID(*service.CategoryID)
		if err != nil {
			return errors.New("invalid category")
		}
	}

	// Update fields
	existing.CategoryID = service.CategoryID
	existing.Name = service.Name
	existing.Description = service.Description
	existing.Price = service.Price
	existing.ImageURL = service.ImageURL
	existing.Active = service.Active

	return s.serviceRepo.Update(existing)
}

func (s *serviceService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid service ID")
	}

	// Check if service exists
	_, err := s.serviceRepo.GetByID(id)
	if err != nil {
		return errors.New("service not found")
	}

	return s.serviceRepo.Delete(id)
}

func (s *serviceService) List() ([]*models.Service, error) {
	return s.serviceRepo.List()
}

func (s *serviceService) ListActive() ([]*models.Service, error) {
	return s.serviceRepo.ListActive()
}

func (s *serviceService) ListByCategory(categoryID int) ([]*models.Service, error) {
	if categoryID <= 0 {
		return nil, errors.New("invalid category ID")
	}

	// Check if category exists
	_, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	return s.serviceRepo.ListByCategory(categoryID)
}
