package services

import (
	"appointment-api/internal/config"
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(req *models.LoginRequest) (*models.AuthResponse, error)
	ValidateToken(tokenString string) (*models.User, error)
	ChangePassword(userID int, currentPassword, newPassword string) error
	UpdateProfile(user *models.User) (*models.User, error)
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
}

type authService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		config:   cfg,
	}
}

func (s *authService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user exists
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if existingUser != nil && existingUser.ID > 0 {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
		Name:     req.Name,
		Phone:    req.Phone,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) ValidateToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var userID int
		switch v := claims["user_id"].(type) {
		case float64:
			userID = int(v)
		case int:
			userID = v
		default:
			return nil, errors.New("invalid user_id in token")
		}

		return s.userRepo.GetByID(userID)
	}

	return nil, errors.New("invalid token")
}

func (s *authService) ChangePassword(userID int, currentPassword, newPassword string) error {
	// Get user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Check current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentPassword)); err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password
	return s.userRepo.UpdatePassword(userID, string(hashedPassword))
}

func (s *authService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWT.Secret))
}

func (s *authService) UpdateProfile(user *models.User) (*models.User, error) {
	// Get existing user to preserve sensitive fields
	existing, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return nil, err
	}

	// Update only allowed fields
	existing.Name = user.Name
	existing.Phone = user.Phone
	existing.BirthDate = user.BirthDate
	existing.UstBel = user.UstBel
	existing.OrtaBel = user.OrtaBel
	existing.AltBel = user.AltBel

	if err := s.userRepo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *authService) ForgotPassword(email string) error {
	// Check if user exists
	_, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	// TODO: Implement actual password reset token generation and email sending
	// For now just return success
	return nil
}

func (s *authService) ResetPassword(token, newPassword string) error {
	// TODO: Implement actual token validation and password reset
	// For now return error indicating not implemented
	return errors.New("password reset functionality not fully implemented")
}
