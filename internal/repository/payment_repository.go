package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"time"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id int) (*models.Payment, error)
	GetByAppointmentID(appointmentID int) (*models.Payment, error)
	GetByUserID(userID int, limit, offset int) ([]*models.Payment, error)
	List(limit, offset int) ([]*models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id int) error
	GetTotalByDateRange(startDate, endDate time.Time) (float64, error)
	GetByStatus(status models.PaymentStatus, limit, offset int) ([]*models.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	query := `
		INSERT INTO payments (appointment_id, device_id, amount, payment_method, transaction_id, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, created_at`

	err := r.db.QueryRow(query,
		payment.AppointmentID,
		payment.DeviceID,
		payment.Amount,
		payment.PaymentMethod,
		payment.TransactionID,
		payment.Status,
	).Scan(&payment.ID, &payment.CreatedAt)

	return err
}

func (r *paymentRepository) GetByID(id int) (*models.Payment, error) {
	payment := &models.Payment{}
	query := `
		SELECT id, appointment_id, device_id, amount, payment_method, transaction_id, status, created_at
		FROM payments WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&payment.ID,
		&payment.AppointmentID,
		&payment.DeviceID,
		&payment.Amount,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&payment.Status,
		&payment.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return payment, err
}

func (r *paymentRepository) GetByAppointmentID(appointmentID int) (*models.Payment, error) {
	payment := &models.Payment{}
	query := `
		SELECT id, appointment_id, device_id, amount, payment_method, transaction_id, status, created_at
		FROM payments WHERE appointment_id = $1`

	err := r.db.QueryRow(query, appointmentID).Scan(
		&payment.ID,
		&payment.AppointmentID,
		&payment.DeviceID,
		&payment.Amount,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&payment.Status,
		&payment.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return payment, err
}

func (r *paymentRepository) GetByUserID(userID int, limit, offset int) ([]*models.Payment, error) {
	query := `
		SELECT p.id, p.appointment_id, p.device_id, p.amount, p.payment_method, p.transaction_id, p.status, p.created_at
		FROM payments p
		JOIN appointments a ON p.appointment_id = a.id
		WHERE a.user_id = $1
		ORDER BY p.created_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := &models.Payment{}
		err := rows.Scan(
			&payment.ID,
			&payment.AppointmentID,
			&payment.DeviceID,
			&payment.Amount,
			&payment.PaymentMethod,
			&payment.TransactionID,
			&payment.Status,
			&payment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) List(limit, offset int) ([]*models.Payment, error) {
	query := `
		SELECT id, appointment_id, device_id, amount, payment_method, transaction_id, status, created_at
		FROM payments
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := &models.Payment{}
		err := rows.Scan(
			&payment.ID,
			&payment.AppointmentID,
			&payment.DeviceID,
			&payment.Amount,
			&payment.PaymentMethod,
			&payment.TransactionID,
			&payment.Status,
			&payment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) Update(payment *models.Payment) error {
	query := `
		UPDATE payments 
		SET amount = $1, payment_method = $2, status = $3
		WHERE id = $4`

	_, err := r.db.Exec(query,
		payment.Amount,
		payment.PaymentMethod,
		payment.Status,
		payment.ID,
	)

	return err
}

func (r *paymentRepository) Delete(id int) error {
	query := `DELETE FROM payments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *paymentRepository) GetTotalByDateRange(startDate, endDate time.Time) (float64, error) {
	var total float64
	query := `
		SELECT COALESCE(SUM(amount), 0) 
		FROM payments 
		WHERE status = $1 AND created_at BETWEEN $2 AND $3`

	err := r.db.QueryRow(query, models.PaymentCompleted, startDate, endDate).Scan(&total)
	return total, err
}

func (r *paymentRepository) GetByStatus(status models.PaymentStatus, limit, offset int) ([]*models.Payment, error) {
	query := `
		SELECT id, appointment_id, device_id, amount, payment_method, transaction_id, status, created_at
		FROM payments 
		WHERE status = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := &models.Payment{}
		err := rows.Scan(
			&payment.ID,
			&payment.AppointmentID,
			&payment.DeviceID,
			&payment.Amount,
			&payment.PaymentMethod,
			&payment.TransactionID,
			&payment.Status,
			&payment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
