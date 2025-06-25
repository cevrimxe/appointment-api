package models

import (
	"time"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID        int        `json:"id" db:"id"`
	Email     string     `json:"email" db:"email" validate:"required,email"`
	Password  string     `json:"-" db:"password" validate:"required,min=6"`
	Role      UserRole   `json:"role" db:"role"`
	Name      string     `json:"name" db:"name" validate:"required"`
	Phone     string     `json:"phone" db:"phone"`
	BirthDate *time.Time `json:"birth_date" db:"birth_date"`
	UstBel    *float64   `json:"ust_bel" db:"ust_bel"`
	OrtaBel   *float64   `json:"orta_bel" db:"orta_bel"`
	AltBel    *float64   `json:"alt_bel" db:"alt_bel"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
