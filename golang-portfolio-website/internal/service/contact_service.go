package service

import (
	"context"
	"portfolio/internal/dto"
	"portfolio/internal/model"
	"portfolio/internal/repository"
)

type ContactService struct {
	repo repository.ContactRepository
}

func NewContactService(repo repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) CreateContact(ctx context.Context, req *dto.CreateContactRequest) (*dto.ContactResponse, error) {
	contact := &model.Contact{
		Name:    req.Name,
		Email:   req.Email,
		Subject: req.Subject,
		Message: req.Message,
	}

	if err := s.repo.Create(ctx, contact); err != nil {
		return nil, err
	}

	return s.toResponse(contact), nil
}

func (s *ContactService) GetAllContacts(ctx context.Context) ([]*dto.ContactResponse, error) {
	contacts, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.ContactResponse
	for _, contact := range contacts {
		responses = append(responses, s.toResponse(contact))
	}
	return responses, nil
}

func (s *ContactService) toResponse(contact *model.Contact) *dto.ContactResponse {
	return &dto.ContactResponse{
		ID:        contact.ID,
		Name:      contact.Name,
		Email:     contact.Email,
		Subject:   contact.Subject,
		Message:   contact.Message,
		IsRead:    contact.IsRead,
		CreatedAt: contact.CreatedAt,
	}
}
