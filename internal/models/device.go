package models

import (
	"time"
)

type Device struct {
	ID         int       `json:"id" db:"id"`
	Brand      string    `json:"brand" db:"brand" validate:"required"`
	Name       string    `json:"name" db:"name" validate:"required"`
	DeviceDate time.Time `json:"device_date" db:"device_date" validate:"required"`
	Price      float64   `json:"price" db:"price" validate:"required,min=0"`
	Active     bool      `json:"active" db:"active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
