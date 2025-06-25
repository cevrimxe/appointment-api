package models

import (
	"time"
)

type AppointmentStatus string
type PaymentStatus string

const (
	StatusPending   AppointmentStatus = "pending"
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCompleted AppointmentStatus = "completed"
	StatusCancelled AppointmentStatus = "cancelled"

	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
)

type Appointment struct {
	ID              int               `json:"id" db:"id"`
	UserID          int               `json:"user_id" db:"user_id"`
	SpecialistID    int               `json:"specialist_id" db:"specialist_id"`
	ServiceID       int               `json:"service_id" db:"service_id"`
	AppointmentDate time.Time         `json:"appointment_date" db:"appointment_date"`
	AppointmentTime time.Time         `json:"appointment_time" db:"appointment_time"`
	Status          AppointmentStatus `json:"status" db:"status"`
	PaymentStatus   PaymentStatus     `json:"payment_status" db:"payment_status"`
	TotalAmount     float64           `json:"total_amount" db:"total_amount"`
	Notes           string            `json:"notes" db:"notes"`
	CreatedAt       time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at" db:"updated_at"`
}

type CreateAppointmentRequest struct {
	SpecialistID    int       `json:"specialist_id" validate:"required"`
	ServiceID       int       `json:"service_id" validate:"required"`
	AppointmentDate time.Time `json:"appointment_date" validate:"required"`
	AppointmentTime time.Time `json:"appointment_time" validate:"required"`
	Notes           string    `json:"notes"`
}

type Service struct {
	ID          int       `json:"id" db:"id"`
	CategoryID  *int      `json:"category_id" db:"category_id"`
	Name        string    `json:"name" db:"name" validate:"required"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price" validate:"required,min=0"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	Active      bool      `json:"active" db:"active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Specialist struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Phone     string    `json:"phone" db:"phone"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type WorkingHour struct {
	ID           int    `json:"id" db:"id"`
	SpecialistID int    `json:"specialist_id" db:"specialist_id"`
	DayOfWeek    int    `json:"day_of_week" db:"day_of_week"` // 0=Sunday, 1=Monday, ...
	StartTime    string `json:"start_time" db:"start_time"`   // HH:MM format
	EndTime      string `json:"end_time" db:"end_time"`       // HH:MM format
	Active       bool   `json:"active" db:"active"`           // Whether this working day is active
}
