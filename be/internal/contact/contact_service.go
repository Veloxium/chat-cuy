package contact

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) GetContactByID(c context.Context, contactID int64) (*GetContactsRes, error) {

}

func (s *service) GetAllContacts(c context.Context) ([]GetContactsRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	contacts, err := s.Repository.GetAllContacts(ctx)
	if err != nil {
		return nil, err
	}

	return contacts, nil

}

func (s *service) DeleteContact(c context.Context, contactID int64) error {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	err := s.Repository.DeleteContact(ctx, contactID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddContact(c context.Context, req *CreateContactReq) (*CreateContactRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	contact := &Contact{
		UserId:   req.UserId,
		Username: req.Username,
		Avatar:   req.Avatar,
	}

	data, err := s.Repository.AddContact(ctx, contact)
	if err != nil {
		return nil, err
	}
	result := &CreateContactRes{
		ID:        data.ID,
		UserId:    data.UserId,
		Avatar:    data.Avatar,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
	return result, nil
}
