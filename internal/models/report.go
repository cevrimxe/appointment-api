package models

import (
	"time"
)

type ReportType string
type ReportStatus string

const (
	ReportTypeSales        ReportType = "sales"
	ReportTypePayments     ReportType = "payments"
	ReportTypeAppointments ReportType = "appointments"
	ReportTypeUsers        ReportType = "users"

	ReportStatusGenerated  ReportStatus = "generated"
	ReportStatusDownloaded ReportStatus = "downloaded"
	ReportStatusExpired    ReportStatus = "expired"
)

type Report struct {
	ID         int          `json:"id" db:"id"`
	UserID     int          `json:"user_id" db:"user_id"`
	ReportType ReportType   `json:"report_type" db:"report_type"`
	FileName   string       `json:"file_name" db:"file_name"`
	FilePath   string       `json:"file_path" db:"file_path"`
	Filters    string       `json:"filters" db:"filters"` // JSON string
	Status     ReportStatus `json:"status" db:"status"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	ExpiresAt  *time.Time   `json:"expires_at" db:"expires_at"`
}

type ReportFilters struct {
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	SpecialistID *int       `json:"specialist_id"`
	ServiceID    *int       `json:"service_id"`
	CategoryID   *int       `json:"category_id"`
	UserID       *int       `json:"user_id"`
	Status       *string    `json:"status"`
}
