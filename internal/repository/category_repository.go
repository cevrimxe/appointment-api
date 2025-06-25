package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"time"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int) error
	List() ([]*models.Category, error)
	ListActive() ([]*models.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	query := `
		INSERT INTO categories (name, description, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(query, category.Name, category.Description,
		category.Active, now, now).Scan(&category.ID)
	if err != nil {
		return err
	}

	category.CreatedAt = now
	category.UpdatedAt = now
	return nil
}

func (r *categoryRepository) GetByID(id int) (*models.Category, error) {
	query := `
		SELECT id, name, description, active, created_at, updated_at
		FROM categories WHERE id = $1`

	category := &models.Category{}
	err := r.db.QueryRow(query, id).Scan(
		&category.ID, &category.Name, &category.Description,
		&category.Active, &category.CreatedAt, &category.UpdatedAt,
	)
	return category, err
}

func (r *categoryRepository) Update(category *models.Category) error {
	query := `
		UPDATE categories 
		SET name = $1, description = $2, active = $3, updated_at = $4
		WHERE id = $5`

	category.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, category.Name, category.Description,
		category.Active, category.UpdatedAt, category.ID)
	return err
}

func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *categoryRepository) List() ([]*models.Category, error) {
	query := `
		SELECT id, name, description, active, created_at, updated_at
		FROM categories
		ORDER BY name ASC`

	return r.queryCategories(query)
}

func (r *categoryRepository) ListActive() ([]*models.Category, error) {
	query := `
		SELECT id, name, description, active, created_at, updated_at
		FROM categories
		WHERE active = true
		ORDER BY name ASC`

	return r.queryCategories(query)
}

func (r *categoryRepository) queryCategories(query string) ([]*models.Category, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		category := &models.Category{}
		err := rows.Scan(
			&category.ID, &category.Name, &category.Description,
			&category.Active, &category.CreatedAt, &category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
