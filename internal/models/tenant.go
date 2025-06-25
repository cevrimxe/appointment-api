package models

import (
	"time"
)

type Tenant struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name" validate:"required"`
	Subdomain  string    `json:"subdomain" db:"subdomain" validate:"required"`
	Domain     string    `json:"domain" db:"domain"`
	SchemaName string    `json:"schema_name" db:"schema_name" validate:"required"`
	Active     bool      `json:"active" db:"active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type TenantConfig struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
	Host   string `json:"host"`
}
