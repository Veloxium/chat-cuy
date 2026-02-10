package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
)

type ContactService interface {
	AddContact(ctx context.Context, userID int, contactUsername string) error
	GetContacts(ctx context.Context, userID int) ([]model.ContactWithUser, error)
}

type contactService struct {
	contactRepo repository.ContactRepository
	userRepo    repository.UserRepository
}

func NewContactService(contactRepo repository.ContactRepository, userRepo repository.UserRepository) ContactService {
	return &contactService{
		contactRepo: contactRepo,
		userRepo:    userRepo,
	}
}

func (s *contactService) AddContact(ctx context.Context, userID int, contactUsername string) error {
	contactUser, err := s.userRepo.GetByUsername(ctx, contactUsername)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	if contactUser.ID == userID {
		return errors.New("cannot add yourself as a contact")
	}

	exists, err := s.contactRepo.Exists(ctx, userID, contactUser.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("contact already added")
	}

	return s.contactRepo.Add(ctx, userID, contactUser.ID)
}

func (s *contactService) GetContacts(ctx context.Context, userID int) ([]model.ContactWithUser, error) {
	return s.contactRepo.GetByUserID(ctx, userID)
}
