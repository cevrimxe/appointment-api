package repository

import (
	"appointment-api/internal/models"
	"database/sql"
	"fmt"
)

type ContactRepository interface {
	Create(message *models.ContactMessage) error
	GetByID(id int) (*models.ContactMessage, error)
	List(limit, offset int) ([]*models.ContactMessage, int, error)
	MarkAsRead(id int) error
	Delete(id int) error
	GetUnreadCount() (int, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(message *models.ContactMessage) error {
	query := `
		INSERT INTO contact_messages (name, email, subject, message, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at`

	err := r.db.QueryRow(query, message.Name, message.Email, message.Subject, message.Message).
		Scan(&message.ID, &message.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create contact message: %v", err)
	}

	return nil
}

func (r *contactRepository) GetByID(id int) (*models.ContactMessage, error) {
	message := &models.ContactMessage{}
	query := `
		SELECT id, name, email, subject, message, is_read, created_at
		FROM contact_messages 
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&message.ID, &message.Name, &message.Email, &message.Subject,
		&message.Message, &message.IsRead, &message.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("contact message not found")
		}
		return nil, fmt.Errorf("failed to get contact message: %v", err)
	}

	return message, nil
}

func (r *contactRepository) List(limit, offset int) ([]*models.ContactMessage, int, error) {
	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM contact_messages"
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %v", err)
	}

	// Get messages with pagination
	query := `
		SELECT id, name, email, subject, message, is_read, created_at
		FROM contact_messages 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list contact messages: %v", err)
	}
	defer rows.Close()

	var messages []*models.ContactMessage
	for rows.Next() {
		message := &models.ContactMessage{}
		err := rows.Scan(
			&message.ID, &message.Name, &message.Email, &message.Subject,
			&message.Message, &message.IsRead, &message.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan contact message: %v", err)
		}
		messages = append(messages, message)
	}

	return messages, total, nil
}

func (r *contactRepository) MarkAsRead(id int) error {
	query := "UPDATE contact_messages SET is_read = true WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to mark message as read: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contact message not found")
	}

	return nil
}

func (r *contactRepository) Delete(id int) error {
	query := "DELETE FROM contact_messages WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete contact message: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contact message not found")
	}

	return nil
}

func (r *contactRepository) GetUnreadCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM contact_messages WHERE is_read = false"
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get unread count: %v", err)
	}
	return count, nil
}
