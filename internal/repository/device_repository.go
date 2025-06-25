package repository

import (
	"appointment-api/internal/models"
	"database/sql"
)

type DeviceRepository interface {
	Create(device *models.Device) error
	GetByID(id int) (*models.Device, error)
	Update(device *models.Device) error
	Delete(id int) error
	List() ([]*models.Device, error)
	ListActive() ([]*models.Device, error)
}

type deviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) DeviceRepository {
	return &deviceRepository{db: db}
}

func (r *deviceRepository) Create(device *models.Device) error {
	query := `
		INSERT INTO devices (brand, name, device_date, price, active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		device.Brand,
		device.Name,
		device.DeviceDate,
		device.Price,
		device.Active,
	).Scan(&device.ID, &device.CreatedAt, &device.UpdatedAt)

	return err
}

func (r *deviceRepository) GetByID(id int) (*models.Device, error) {
	query := `
		SELECT id, brand, name, device_date, price, active, created_at, updated_at
		FROM devices
		WHERE id = $1`

	device := &models.Device{}
	err := r.db.QueryRow(query, id).Scan(
		&device.ID,
		&device.Brand,
		&device.Name,
		&device.DeviceDate,
		&device.Price,
		&device.Active,
		&device.CreatedAt,
		&device.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return device, nil
}

func (r *deviceRepository) Update(device *models.Device) error {
	query := `
		UPDATE devices 
		SET brand = $2, name = $3, device_date = $4, price = $5, active = $6, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		device.ID,
		device.Brand,
		device.Name,
		device.DeviceDate,
		device.Price,
		device.Active,
	).Scan(&device.UpdatedAt)

	return err
}

func (r *deviceRepository) Delete(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *deviceRepository) List() ([]*models.Device, error) {
	query := `
		SELECT id, brand, name, device_date, price, active, created_at, updated_at
		FROM devices
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*models.Device
	for rows.Next() {
		device := &models.Device{}
		err := rows.Scan(
			&device.ID,
			&device.Brand,
			&device.Name,
			&device.DeviceDate,
			&device.Price,
			&device.Active,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func (r *deviceRepository) ListActive() ([]*models.Device, error) {
	query := `
		SELECT id, brand, name, device_date, price, active, created_at, updated_at
		FROM devices
		WHERE active = true
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []*models.Device
	for rows.Next() {
		device := &models.Device{}
		err := rows.Scan(
			&device.ID,
			&device.Brand,
			&device.Name,
			&device.DeviceDate,
			&device.Price,
			&device.Active,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}
