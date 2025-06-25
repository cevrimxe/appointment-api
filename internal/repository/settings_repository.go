package repository

import (
	"appointment-api/internal/models"
	"database/sql"
)

type SettingsRepository interface {
	GetByKey(key string) (*models.Setting, error)
	UpdateByKey(key, value, description string) error
	List() ([]*models.Setting, error)
}

type settingsRepository struct {
	db *sql.DB
}

func NewSettingsRepository(db *sql.DB) SettingsRepository {
	return &settingsRepository{db: db}
}

func (r *settingsRepository) GetByKey(key string) (*models.Setting, error) {
	query := `
		SELECT id, key, value, description, created_at, updated_at
		FROM settings WHERE key = $1`

	setting := &models.Setting{}
	err := r.db.QueryRow(query, key).Scan(
		&setting.ID, &setting.Key, &setting.Value,
		&setting.Description, &setting.CreatedAt, &setting.UpdatedAt,
	)
	return setting, err
}

func (r *settingsRepository) UpdateByKey(key, value, description string) error {
	query := `
		UPDATE settings 
		SET value = $1, description = $2, updated_at = CURRENT_TIMESTAMP
		WHERE key = $3`

	_, err := r.db.Exec(query, value, description, key)
	return err
}

func (r *settingsRepository) List() ([]*models.Setting, error) {
	query := `
		SELECT id, key, value, description, created_at, updated_at
		FROM settings
		ORDER BY key ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settings []*models.Setting
	for rows.Next() {
		setting := &models.Setting{}
		err := rows.Scan(
			&setting.ID, &setting.Key, &setting.Value,
			&setting.Description, &setting.CreatedAt, &setting.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		settings = append(settings, setting)
	}

	return settings, nil
}
