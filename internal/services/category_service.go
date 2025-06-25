package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
)

type CategoryService interface {
	Create(category *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
	List() ([]*models.Category, error)
	ListActive() ([]*models.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) Create(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Set default active status
	category.Active = true

	return s.categoryRepo.Create(category)
}

func (s *categoryService) GetByID(id int) (*models.Category, error) {
	if id <= 0 {
		return nil, errors.New("invalid category ID")
	}

	return s.categoryRepo.GetByID(id)
}

func (s *categoryService) Update(category *models.Category) error {
	if category.ID <= 0 {
		return errors.New("invalid category ID")
	}

	if category.Name == "" {
		return errors.New("category name is required")
	}

	// Check if category exists
	existing, err := s.categoryRepo.GetByID(category.ID)
	if err != nil {
		return errors.New("category not found")
	}

	// Update fields
	existing.Name = category.Name
	existing.Description = category.Description
	existing.Active = category.Active

	return s.categoryRepo.Update(existing)
}

func (s *categoryService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid category ID")
	}

	// Check if category exists
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	return s.categoryRepo.Delete(id)
}

func (s *categoryService) List() ([]*models.Category, error) {
	return s.categoryRepo.List()
}

func (s *categoryService) ListActive() ([]*models.Category, error) {
	return s.categoryRepo.ListActive()
}
