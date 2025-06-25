package repository

import (
	"appointment-api/internal/models"
	"database/sql"
)

type SpecialistRepository interface {
	Create(specialist *models.Specialist) error
	GetByID(id int) (*models.Specialist, error)
	Update(specialist *models.Specialist) error
	Delete(id int) error
	List() ([]*models.Specialist, error)
	ListActive() ([]*models.Specialist, error)
	GetWorkingHours(specialistID int) ([]*models.WorkingHour, error)
	UpdateWorkingHours(specialistID int, workingHours []*models.WorkingHour) error
}

type specialistRepository struct {
	db *sql.DB
}

func NewSpecialistRepository(db *sql.DB) SpecialistRepository {
	return &specialistRepository{db: db}
}

func (r *specialistRepository) Create(specialist *models.Specialist) error {
	query := `
		INSERT INTO specialists (name, email, phone, active)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		specialist.Name,
		specialist.Email,
		specialist.Phone,
		specialist.Active,
	).Scan(&specialist.ID, &specialist.CreatedAt, &specialist.UpdatedAt)

	return err
}

func (r *specialistRepository) GetByID(id int) (*models.Specialist, error) {
	query := `
		SELECT id, name, email, phone, active, created_at, updated_at
		FROM specialists
		WHERE id = $1`

	specialist := &models.Specialist{}
	err := r.db.QueryRow(query, id).Scan(
		&specialist.ID,
		&specialist.Name,
		&specialist.Email,
		&specialist.Phone,
		&specialist.Active,
		&specialist.CreatedAt,
		&specialist.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return specialist, nil
}

func (r *specialistRepository) Update(specialist *models.Specialist) error {
	query := `
		UPDATE specialists 
		SET name = $2, email = $3, phone = $4, active = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING updated_at`

	err := r.db.QueryRow(
		query,
		specialist.ID,
		specialist.Name,
		specialist.Email,
		specialist.Phone,
		specialist.Active,
	).Scan(&specialist.UpdatedAt)

	return err
}

func (r *specialistRepository) Delete(id int) error {
	query := `DELETE FROM specialists WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *specialistRepository) List() ([]*models.Specialist, error) {
	query := `
		SELECT id, name, email, phone, active, created_at, updated_at
		FROM specialists
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialists []*models.Specialist
	for rows.Next() {
		specialist := &models.Specialist{}
		err := rows.Scan(
			&specialist.ID,
			&specialist.Name,
			&specialist.Email,
			&specialist.Phone,
			&specialist.Active,
			&specialist.CreatedAt,
			&specialist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		specialists = append(specialists, specialist)
	}

	return specialists, nil
}

func (r *specialistRepository) ListActive() ([]*models.Specialist, error) {
	query := `
		SELECT id, name, email, phone, active, created_at, updated_at
		FROM specialists
		WHERE active = true
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialists []*models.Specialist
	for rows.Next() {
		specialist := &models.Specialist{}
		err := rows.Scan(
			&specialist.ID,
			&specialist.Name,
			&specialist.Email,
			&specialist.Phone,
			&specialist.Active,
			&specialist.CreatedAt,
			&specialist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		specialists = append(specialists, specialist)
	}

	return specialists, nil
}

func (r *specialistRepository) GetWorkingHours(specialistID int) ([]*models.WorkingHour, error) {
	query := `
		SELECT id, specialist_id, day_of_week, start_time, end_time, active
		FROM working_hours
		WHERE specialist_id = $1
		ORDER BY day_of_week`

	rows, err := r.db.Query(query, specialistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workingHours []*models.WorkingHour
	for rows.Next() {
		wh := &models.WorkingHour{}
		err := rows.Scan(
			&wh.ID,
			&wh.SpecialistID,
			&wh.DayOfWeek,
			&wh.StartTime,
			&wh.EndTime,
			&wh.Active,
		)
		if err != nil {
			return nil, err
		}
		workingHours = append(workingHours, wh)
	}

	return workingHours, nil
}

func (r *specialistRepository) UpdateWorkingHours(specialistID int, workingHours []*models.WorkingHour) error {
	// Start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Delete existing working hours
	_, err = tx.Exec("DELETE FROM working_hours WHERE specialist_id = $1", specialistID)
	if err != nil {
		return err
	}

	// Insert new working hours
	for _, wh := range workingHours {
		query := `
			INSERT INTO working_hours (specialist_id, day_of_week, start_time, end_time, active)
			VALUES ($1, $2, $3, $4, $5)`

		_, err = tx.Exec(query, specialistID, wh.DayOfWeek, wh.StartTime, wh.EndTime, wh.Active)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
