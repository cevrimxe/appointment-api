package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"time"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	GetByID(id int) (*models.Appointment, error)
	Update(appointment *models.Appointment) error
	Delete(id int) error
	GetByUserID(userID int) ([]*models.Appointment, error)
	GetBySpecialistID(specialistID int, date *time.Time) ([]*models.Appointment, error)
	List(limit, offset int) ([]*models.Appointment, int, error)
	UpdateStatus(id int, status models.AppointmentStatus) error
	CheckConflict(specialistID int, appointmentDate, appointmentTime time.Time, excludeID *int) (bool, error)
	UpdatePaymentStatus(appointmentID int, status models.PaymentStatus) error
}

type appointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(appointment *models.Appointment) error {
	query := `
		INSERT INTO appointments (user_id, specialist_id, service_id, appointment_date, appointment_time, 
			status, payment_status, total_amount, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		appointment.UserID,
		appointment.SpecialistID,
		appointment.ServiceID,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Status,
		appointment.PaymentStatus,
		appointment.TotalAmount,
		appointment.Notes,
		now,
		now,
	).Scan(&appointment.ID)

	if err != nil {
		return err
	}

	appointment.CreatedAt = now
	appointment.UpdatedAt = now
	return nil
}

func (r *appointmentRepository) GetByID(id int) (*models.Appointment, error) {
	query := `
		SELECT id, user_id, specialist_id, service_id, appointment_date, appointment_time,
			status, payment_status, total_amount, notes, created_at, updated_at
		FROM appointments 
		WHERE id = $1`

	appointment := &models.Appointment{}
	err := r.db.QueryRow(query, id).Scan(
		&appointment.ID,
		&appointment.UserID,
		&appointment.SpecialistID,
		&appointment.ServiceID,
		&appointment.AppointmentDate,
		&appointment.AppointmentTime,
		&appointment.Status,
		&appointment.PaymentStatus,
		&appointment.TotalAmount,
		&appointment.Notes,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)

	return appointment, err
}

func (r *appointmentRepository) Update(appointment *models.Appointment) error {
	query := `
		UPDATE appointments 
		SET specialist_id = $2, service_id = $3, appointment_date = $4, appointment_time = $5,
			status = $6, payment_status = $7, total_amount = $8, notes = $9, updated_at = $10
		WHERE id = $1
		RETURNING updated_at`

	appointment.UpdatedAt = time.Now()
	err := r.db.QueryRow(
		query,
		appointment.ID,
		appointment.SpecialistID,
		appointment.ServiceID,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Status,
		appointment.PaymentStatus,
		appointment.TotalAmount,
		appointment.Notes,
		appointment.UpdatedAt,
	).Scan(&appointment.UpdatedAt)

	return err
}

func (r *appointmentRepository) Delete(id int) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *appointmentRepository) GetByUserID(userID int) ([]*models.Appointment, error) {
	query := `
		SELECT id, user_id, specialist_id, service_id, appointment_date, appointment_time,
			status, payment_status, total_amount, notes, created_at, updated_at
		FROM appointments 
		WHERE user_id = $1
		ORDER BY appointment_date DESC, appointment_time DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*models.Appointment
	for rows.Next() {
		appointment := &models.Appointment{}
		err := rows.Scan(
			&appointment.ID,
			&appointment.UserID,
			&appointment.SpecialistID,
			&appointment.ServiceID,
			&appointment.AppointmentDate,
			&appointment.AppointmentTime,
			&appointment.Status,
			&appointment.PaymentStatus,
			&appointment.TotalAmount,
			&appointment.Notes,
			&appointment.CreatedAt,
			&appointment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (r *appointmentRepository) GetBySpecialistID(specialistID int, date *time.Time) ([]*models.Appointment, error) {
	var query string
	var args []interface{}

	if date != nil {
		query = `
			SELECT id, user_id, specialist_id, service_id, appointment_date, appointment_time,
				status, payment_status, total_amount, notes, created_at, updated_at
			FROM appointments 
			WHERE specialist_id = $1 AND appointment_date = $2
			ORDER BY appointment_time ASC`
		args = []interface{}{specialistID, *date}
	} else {
		query = `
			SELECT id, user_id, specialist_id, service_id, appointment_date, appointment_time,
				status, payment_status, total_amount, notes, created_at, updated_at
			FROM appointments 
			WHERE specialist_id = $1
			ORDER BY appointment_date DESC, appointment_time DESC`
		args = []interface{}{specialistID}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*models.Appointment
	for rows.Next() {
		appointment := &models.Appointment{}
		err := rows.Scan(
			&appointment.ID,
			&appointment.UserID,
			&appointment.SpecialistID,
			&appointment.ServiceID,
			&appointment.AppointmentDate,
			&appointment.AppointmentTime,
			&appointment.Status,
			&appointment.PaymentStatus,
			&appointment.TotalAmount,
			&appointment.Notes,
			&appointment.CreatedAt,
			&appointment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (r *appointmentRepository) List(limit, offset int) ([]*models.Appointment, int, error) {
	// Count total
	countQuery := `SELECT COUNT(*) FROM appointments`
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get appointments
	query := `
		SELECT id, user_id, specialist_id, service_id, appointment_date, appointment_time,
			status, payment_status, total_amount, notes, created_at, updated_at
		FROM appointments 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var appointments []*models.Appointment
	for rows.Next() {
		appointment := &models.Appointment{}
		err := rows.Scan(
			&appointment.ID,
			&appointment.UserID,
			&appointment.SpecialistID,
			&appointment.ServiceID,
			&appointment.AppointmentDate,
			&appointment.AppointmentTime,
			&appointment.Status,
			&appointment.PaymentStatus,
			&appointment.TotalAmount,
			&appointment.Notes,
			&appointment.CreatedAt,
			&appointment.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, total, nil
}

func (r *appointmentRepository) UpdateStatus(id int, status models.AppointmentStatus) error {
	query := `
		UPDATE appointments 
		SET status = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1`

	_, err := r.db.Exec(query, id, status)
	return err
}

func (r *appointmentRepository) CheckConflict(specialistID int, appointmentDate, appointmentTime time.Time, excludeID *int) (bool, error) {
	var query string
	var args []interface{}

	if excludeID != nil {
		query = `
			SELECT COUNT(*) FROM appointments 
			WHERE specialist_id = $1 AND appointment_date = $2 AND appointment_time = $3 
			AND status != 'cancelled' AND id != $4`
		args = []interface{}{specialistID, appointmentDate, appointmentTime, *excludeID}
	} else {
		query = `
			SELECT COUNT(*) FROM appointments 
			WHERE specialist_id = $1 AND appointment_date = $2 AND appointment_time = $3 
			AND status != 'cancelled'`
		args = []interface{}{specialistID, appointmentDate, appointmentTime}
	}

	var count int
	err := r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *appointmentRepository) UpdatePaymentStatus(appointmentID int, status models.PaymentStatus) error {
	query := `UPDATE appointments SET payment_status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, appointmentID)
	return err
}
