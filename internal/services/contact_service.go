package services

import (
	"appointment-api/internal/models"
	"appointment-api/internal/repository"
)

type ContactService interface {
	Create(request *models.ContactMessageRequest) (*models.ContactMessage, error)
	GetByID(id int) (*models.ContactMessage, error)
	List(limit, offset int) ([]*models.ContactMessage, int, error)
	MarkAsRead(id int) error
	Delete(id int) error
	GetUnreadCount() (int, error)
}

type contactService struct {
	contactRepo repository.ContactRepository
}

func NewContactService(contactRepo repository.ContactRepository) ContactService {
	return &contactService{
		contactRepo: contactRepo,
	}
}

func (s *contactService) Create(request *models.ContactMessageRequest) (*models.ContactMessage, error) {
	message := &models.ContactMessage{
		Name:    request.Name,
		Email:   request.Email,
		Subject: request.Subject,
		Message: request.Message,
		IsRead:  false,
	}

	err := s.contactRepo.Create(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *contactService) GetByID(id int) (*models.ContactMessage, error) {
	return s.contactRepo.GetByID(id)
}

func (s *contactService) List(limit, offset int) ([]*models.ContactMessage, int, error) {
	return s.contactRepo.List(limit, offset)
}

func (s *contactService) MarkAsRead(id int) error {
	return s.contactRepo.MarkAsRead(id)
}

func (s *contactService) Delete(id int) error {
	return s.contactRepo.Delete(id)
}

func (s *contactService) GetUnreadCount() (int, error) {
	return s.contactRepo.GetUnreadCount()
}
