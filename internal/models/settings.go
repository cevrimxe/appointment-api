package models

import (
	"time"
)

type Setting struct {
	ID          int       `json:"id" db:"id"`
	Key         string    `json:"key" db:"key" validate:"required"`
	Value       string    `json:"value" db:"value" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
