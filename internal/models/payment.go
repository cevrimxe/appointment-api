package models

import (
	"time"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodCash       PaymentMethod = "cash"
	PaymentMethodTransfer   PaymentMethod = "transfer"
)

type Payment struct {
	ID            int           `json:"id" db:"id"`
	AppointmentID int           `json:"appointment_id" db:"appointment_id"`
	DeviceID      *int          `json:"device_id" db:"device_id"`
	Amount        float64       `json:"amount" db:"amount" validate:"required,min=0"`
	PaymentMethod PaymentMethod `json:"payment_method" db:"payment_method"`
	TransactionID string        `json:"transaction_id" db:"transaction_id"`
	Status        PaymentStatus `json:"status" db:"status"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
}

type CreatePaymentRequest struct {
	AppointmentID int           `json:"appointment_id" validate:"required"`
	Amount        float64       `json:"amount" validate:"required,min=0"`
	PaymentMethod PaymentMethod `json:"payment_method" validate:"required"`
	DeviceID      *int          `json:"device_id"`
}
