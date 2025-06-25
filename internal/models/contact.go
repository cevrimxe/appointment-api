package models

import "time"

type ContactMessage struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Subject   string    `json:"subject" db:"subject" validate:"required"`
	Message   string    `json:"message" db:"message" validate:"required"`
	IsRead    bool      `json:"is_read" db:"is_read"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ContactMessageRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
}
