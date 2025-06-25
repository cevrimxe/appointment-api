package services

import (
	"appointment-api/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type TenantService interface {
	GetTenantByHost(host string) (*models.TenantConfig, error)
	CreateTenantSchema(schemaName string) error
	GetAllTenants() ([]*models.Tenant, error)
}

type tenantService struct {
	db *sql.DB
}

func NewTenantService(db *sql.DB) TenantService {
	return &tenantService{db: db}
}

func (s *tenantService) GetTenantByHost(host string) (*models.TenantConfig, error) {
	var subdomain string

	// Remove port from host if present
	hostParts := strings.Split(host, ":")
	hostWithoutPort := hostParts[0]

	// Extract subdomain from host
	parts := strings.Split(hostWithoutPort, ".")

	if len(parts) == 1 {
		// Single part (e.g., "localhost")
		if hostWithoutPort == "localhost" || hostWithoutPort == "127.0.0.1" {
			subdomain = "main"
		} else {
			return nil, fmt.Errorf("invalid host format")
		}
	} else {
		// Multiple parts (e.g., "test.localhost", "www.example.com")
		subdomain = parts[0]
		if subdomain == "www" && len(parts) > 2 {
			subdomain = parts[1]
		}
	}

	query := `
		SELECT id, name, schema_name, subdomain || '.' || COALESCE(domain, 'localhost:8080') as host
		FROM public.tenants 
		WHERE subdomain = $1 AND active = true`

	tenant := &models.TenantConfig{}
	err := s.db.QueryRow(query, subdomain).Scan(
		&tenant.ID, &tenant.Name, &tenant.Schema, &tenant.Host,
	)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *tenantService) CreateTenantSchema(schemaName string) error {
	// Create schema
	_, err := s.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
	if err != nil {
		return err
	}

	// Create tables in the new schema
	err = s.createTenantTables(schemaName)
	if err != nil {
		// Rollback schema creation
		s.db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schemaName))
		return err
	}

	return nil
}

func (s *tenantService) GetAllTenants() ([]*models.Tenant, error) {
	query := `
		SELECT id, name, subdomain, domain, schema_name, active, created_at, updated_at
		FROM public.tenants
		ORDER BY created_at DESC`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []*models.Tenant
	for rows.Next() {
		tenant := &models.Tenant{}
		err := rows.Scan(
			&tenant.ID, &tenant.Name, &tenant.Subdomain, &tenant.Domain,
			&tenant.SchemaName, &tenant.Active, &tenant.CreatedAt, &tenant.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)
	}

	return tenants, nil
}

func (s *tenantService) createTenantTables(schemaName string) error {
	// Read the complete tenant schema template
	templateSQL, err := s.loadTenantSchemaTemplate()
	if err != nil {
		return fmt.Errorf("failed to load schema template: %v", err)
	}

	// Replace {SCHEMA_NAME} placeholder with actual schema name
	schemaSQL := strings.ReplaceAll(templateSQL, "{SCHEMA_NAME}", schemaName)

	// Execute the complete schema creation
	if _, err := s.db.Exec(schemaSQL); err != nil {
		return fmt.Errorf("failed to create tenant schema: %v", err)
	}

	return nil
}

// loadTenantSchemaTemplate loads the complete tenant schema template
func (s *tenantService) loadTenantSchemaTemplate() (string, error) {
	// For now, return the template inline - in production you might read from file
	return `
-- Users table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user' CHECK (role IN ('admin', 'user')),
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    birth_date DATE,
    ust_bel DECIMAL(5,2),
    orta_bel DECIMAL(5,2),
    alt_bel DECIMAL(5,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Settings table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.settings (
    id SERIAL PRIMARY KEY,
    key VARCHAR(100) UNIQUE NOT NULL,
    value TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Services table (WITHOUT duration_minutes - using global settings)
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.services (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES {SCHEMA_NAME}.categories(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    image_url VARCHAR(500),
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Specialists table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.specialists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Working hours table (WITH active column)
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.working_hours (
    id SERIAL PRIMARY KEY,
    specialist_id INTEGER REFERENCES {SCHEMA_NAME}.specialists(id) ON DELETE CASCADE,
    day_of_week INTEGER NOT NULL CHECK (day_of_week >= 0 AND day_of_week <= 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    active BOOLEAN DEFAULT true,
    UNIQUE(specialist_id, day_of_week)
);

-- Devices table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.devices (
    id SERIAL PRIMARY KEY,
    brand VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    device_date DATE NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Appointments table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.appointments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES {SCHEMA_NAME}.users(id) ON DELETE CASCADE,
    specialist_id INTEGER REFERENCES {SCHEMA_NAME}.specialists(id) ON DELETE CASCADE,
    service_id INTEGER REFERENCES {SCHEMA_NAME}.services(id) ON DELETE CASCADE,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'confirmed', 'completed', 'cancelled')),
    payment_status VARCHAR(20) DEFAULT 'pending' CHECK (payment_status IN ('pending', 'completed', 'failed', 'refunded')),
    total_amount DECIMAL(10,2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(specialist_id, appointment_date, appointment_time)
);

-- Contact messages table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.contact_messages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Payments table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.payments (
    id SERIAL PRIMARY KEY,
    appointment_id INTEGER REFERENCES {SCHEMA_NAME}.appointments(id) ON DELETE CASCADE,
    device_id INTEGER REFERENCES {SCHEMA_NAME}.devices(id) ON DELETE SET NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_method VARCHAR(50),
    transaction_id VARCHAR(255),
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'failed', 'refunded')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reports table
CREATE TABLE IF NOT EXISTS {SCHEMA_NAME}.reports (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES {SCHEMA_NAME}.users(id) ON DELETE CASCADE,
    report_type VARCHAR(50) NOT NULL CHECK (report_type IN ('sales', 'payments', 'appointments', 'users')),
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(500),
    filters JSONB,
    status VARCHAR(20) DEFAULT 'generated' CHECK (status IN ('generated', 'downloaded', 'expired')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_appointments_user_id ON {SCHEMA_NAME}.appointments(user_id);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_appointments_specialist_date ON {SCHEMA_NAME}.appointments(specialist_id, appointment_date);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_appointments_date ON {SCHEMA_NAME}.appointments(appointment_date);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_working_hours_specialist ON {SCHEMA_NAME}.working_hours(specialist_id);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_working_hours_active ON {SCHEMA_NAME}.working_hours(specialist_id, active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_contact_messages_read ON {SCHEMA_NAME}.contact_messages(is_read);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_reports_user_type ON {SCHEMA_NAME}.reports(user_id, report_type);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_reports_created ON {SCHEMA_NAME}.reports(created_at);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_categories_active ON {SCHEMA_NAME}.categories(active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_services_category ON {SCHEMA_NAME}.services(category_id, active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_services_active ON {SCHEMA_NAME}.services(active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_specialists_active ON {SCHEMA_NAME}.specialists(active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_devices_active ON {SCHEMA_NAME}.devices(active);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_payments_appointment ON {SCHEMA_NAME}.payments(appointment_id);
CREATE INDEX IF NOT EXISTS idx_{SCHEMA_NAME}_payments_device ON {SCHEMA_NAME}.payments(device_id);

-- Default settings (including appointment duration for available slots)
INSERT INTO {SCHEMA_NAME}.settings (key, value, description) VALUES 
('appointment_duration', '60', 'Appointment duration in minutes for available slots calculation'),
('working_hours_start', '09:00', 'Default working hours start time'),
('working_hours_end', '17:00', 'Default working hours end time'),
('max_advance_booking_days', '30', 'Maximum days in advance for booking');

-- Sample categories
INSERT INTO {SCHEMA_NAME}.categories (name, description) VALUES 
('Danışmanlık', 'Uzman danışmanlık hizmetleri'),
('Sağlık', 'Sağlık ile ilgili hizmetler'),
('Eğitim', 'Eğitim ve öğretim hizmetleri');

-- Sample services (no duration_minutes - using global setting)
INSERT INTO {SCHEMA_NAME}.services (category_id, name, description, price, image_url) VALUES 
(1, 'Bireysel Danışmanlık', 'Uzman danışmanlık hizmeti', 200.00, 'https://example.com/images/counseling.jpg'),
(1, 'Konsültasyon', 'Kısa konsültasyon hizmeti', 100.00, 'https://example.com/images/consultation.jpg'),
(2, 'Sağlık Kontrolü', 'Detaylı sağlık incelemesi', 300.00, 'https://example.com/images/health-check.jpg');

-- Sample specialist
INSERT INTO {SCHEMA_NAME}.specialists (name, email, phone) VALUES 
('Dr. Ahmet Yılmaz', 'ahmet@example.com', '+90555987654');

-- Sample working hours (with active column - Monday to Friday)
INSERT INTO {SCHEMA_NAME}.working_hours (specialist_id, day_of_week, start_time, end_time, active) VALUES 
(1, 1, '09:00', '17:00', true),
(1, 2, '09:00', '17:00', true),
(1, 3, '09:00', '17:00', true),
(1, 4, '09:00', '17:00', true),
(1, 5, '09:00', '17:00', true);

-- Sample devices
INSERT INTO {SCHEMA_NAME}.devices (brand, name, device_date, price) VALUES 
('Philips', 'X-Ray Cihazı Model A', '2024-01-15', 25000.00),
('Siemens', 'MRI Tarayıcı Pro', '2023-12-20', 85000.00),
('GE Healthcare', 'Ultrason Cihazı', '2024-02-10', 15000.00);
`, nil
}
