package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetByID(id int) (*models.User, error)
	List(limit, offset int) ([]*models.User, int, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
	UpdateRole(userID int, role models.UserRole) error
	GetTotalCount() (int, error)
	GetNewMonthlyCount() (int, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return s.userRepo.GetByID(id)
}

func (s *userService) List(limit, offset int) ([]*models.User, int, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if offset < 0 {
		offset = 0
	}
	return s.userRepo.List(limit, offset)
}

func (s *userService) Create(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = models.RoleUser
	}

	// Validate role
	if user.Role != models.RoleAdmin && user.Role != models.RoleUser {
		return errors.New("invalid role")
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

func (s *userService) Update(user *models.User) error {
	if user.ID <= 0 {
		return errors.New("invalid user ID")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	// Check if user exists
	existing, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	// Check if email is taken by another user
	if user.Email != existing.Email {
		existingByEmail, err := s.userRepo.GetByEmail(user.Email)
		if err == nil && existingByEmail != nil && existingByEmail.ID != user.ID {
			return errors.New("email already exists")
		}
	}

	// Validate role
	if user.Role != "" && user.Role != models.RoleAdmin && user.Role != models.RoleUser {
		return errors.New("invalid role")
	}

	// Update fields but keep password unchanged if not provided
	existing.Email = user.Email
	existing.Name = user.Name
	existing.Phone = user.Phone
	existing.BirthDate = user.BirthDate
	existing.UstBel = user.UstBel
	existing.OrtaBel = user.OrtaBel
	existing.AltBel = user.AltBel

	if user.Role != "" {
		existing.Role = user.Role
	}

	return s.userRepo.Update(existing)
}

func (s *userService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	// Check if user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return s.userRepo.Delete(id)
}

func (s *userService) UpdateRole(userID int, role models.UserRole) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	// Validate role
	if role != models.RoleAdmin && role != models.RoleUser {
		return errors.New("invalid role")
	}

	// Check if user exists
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	user.Role = role
	return s.userRepo.Update(user)
}

func (s *userService) GetTotalCount() (int, error) {
	// Use List method to get total count
	_, total, err := s.userRepo.List(1, 0) // Just get 1 record to get total count
	return total, err
}

func (s *userService) GetNewMonthlyCount() (int, error) {
	// Get all users and filter by this month's creation date
	users, _, err := s.userRepo.List(10000, 0) // Get all users
	if err != nil {
		return 0, err
	}

	now := time.Now()
	currentMonth := now.Month()
	currentYear := now.Year()
	count := 0
	for _, user := range users {
		if user.CreatedAt.Month() == currentMonth && user.CreatedAt.Year() == currentYear {
			count++
		}
	}
	return count, nil
}
