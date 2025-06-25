package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"time"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	Update(user *models.User) error
	UpdatePassword(userID int, hashedPassword string) error
	List(limit, offset int) ([]*models.User, int, error)
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	// Simple query without fixed schema - uses TenantMiddleware's search_path

	query := `
		INSERT INTO users (email, password, role, name, phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	now := time.Now()
	var id int
	err := r.db.QueryRow(query, user.Email, user.Password, user.Role,
		user.Name, user.Phone, now, now).Scan(&id)
	if err != nil {
		return err
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password, role, name, phone, birth_date, 
			   ust_bel, orta_bel, alt_bel, created_at, updated_at
		FROM users WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Role,
		&user.Name, &user.Phone, &user.BirthDate,
		&user.UstBel, &user.OrtaBel, &user.AltBel,
		&user.CreatedAt, &user.UpdatedAt,
	)
	return user, err
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	query := `
		SELECT id, email, password, role, name, phone, birth_date,
			   ust_bel, orta_bel, alt_bel, created_at, updated_at
		FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.Role,
		&user.Name, &user.Phone, &user.BirthDate,
		&user.UstBel, &user.OrtaBel, &user.AltBel,
		&user.CreatedAt, &user.UpdatedAt,
	)
	return user, err
}

func (r *userRepository) Update(user *models.User) error {
	query := `
		UPDATE users 
		SET name = $1, phone = $2, birth_date = $3, ust_bel = $4, 
			orta_bel = $5, alt_bel = $6, updated_at = $7
		WHERE id = $8`

	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, user.Name, user.Phone, user.BirthDate,
		user.UstBel, user.OrtaBel, user.AltBel, user.UpdatedAt, user.ID)
	return err
}

func (r *userRepository) UpdatePassword(userID int, hashedPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, hashedPassword, time.Now(), userID)
	return err
}

func (r *userRepository) List(limit, offset int) ([]*models.User, int, error) {
	// Count total
	var total int
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get users
	query := `
		SELECT id, email, password, role, name, phone, birth_date,
			   ust_bel, orta_bel, alt_bel, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.Password, &user.Role,
			&user.Name, &user.Phone, &user.BirthDate,
			&user.UstBel, &user.OrtaBel, &user.AltBel,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, total, nil
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
