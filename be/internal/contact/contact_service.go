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
