package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"time"
)

type ServiceRepository interface {
	Create(service *models.Service) error
	GetByID(id int) (*models.Service, error)
	Update(service *models.Service) error
	Delete(id int) error
	List() ([]*models.Service, error)
	ListActive() ([]*models.Service, error)
	ListByCategory(categoryID int) ([]*models.Service, error)
}

type serviceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(service *models.Service) error {
	query := `
		INSERT INTO services (category_id, name, description, price, image_url, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(query, service.CategoryID, service.Name, service.Description,
		service.Price, service.ImageURL, service.Active, now, now).Scan(&service.ID)
	if err != nil {
		return err
	}

	service.CreatedAt = now
	service.UpdatedAt = now
	return nil
}

func (r *serviceRepository) GetByID(id int) (*models.Service, error) {
	query := `
		SELECT id, category_id, name, description, price, image_url, active, created_at, updated_at
		FROM services WHERE id = $1`

	service := &models.Service{}
	err := r.db.QueryRow(query, id).Scan(
		&service.ID, &service.CategoryID, &service.Name, &service.Description,
		&service.Price, &service.ImageURL, &service.Active, &service.CreatedAt, &service.UpdatedAt,
	)
	return service, err
}

func (r *serviceRepository) Update(service *models.Service) error {
	query := `
		UPDATE services 
		SET category_id = $1, name = $2, description = $3, price = $4, 
			image_url = $5, active = $6, updated_at = $7
		WHERE id = $8`

	service.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, service.CategoryID, service.Name, service.Description,
		service.Price, service.ImageURL, service.Active, service.UpdatedAt, service.ID)
	return err
}

func (r *serviceRepository) Delete(id int) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *serviceRepository) List() ([]*models.Service, error) {
	query := `
		SELECT id, category_id, name, description, price, image_url, active, created_at, updated_at
		FROM services
		ORDER BY name ASC`

	return r.queryServices(query)
}

func (r *serviceRepository) ListActive() ([]*models.Service, error) {
	query := `
		SELECT id, category_id, name, description, price, image_url, active, created_at, updated_at
		FROM services
		WHERE active = true
		ORDER BY name ASC`

	return r.queryServices(query)
}

func (r *serviceRepository) ListByCategory(categoryID int) ([]*models.Service, error) {
	query := `
		SELECT id, category_id, name, description, price, image_url, active, created_at, updated_at
		FROM services
		WHERE category_id = $1 AND active = true
		ORDER BY name ASC`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanServices(rows)
}

func (r *serviceRepository) queryServices(query string) ([]*models.Service, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanServices(rows)
}

func (r *serviceRepository) scanServices(rows *sql.Rows) ([]*models.Service, error) {
	var services []*models.Service
	for rows.Next() {
		service := &models.Service{}
		err := rows.Scan(
			&service.ID, &service.CategoryID, &service.Name, &service.Description,
			&service.Price, &service.ImageURL, &service.Active, &service.CreatedAt, &service.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}
