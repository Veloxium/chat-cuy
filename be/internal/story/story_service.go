package story

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

// CreateStory implements Service.
func (s *service) CreateStory(ctx context.Context, req *CreateStoryReq) (*CreateStoryRes, error) {
	story := &Story{
		UserID:      req.UserID,
		ContentType: req.ContentType,
		Content:     req.Content,
	}

	res, err := s.Repository.CreateStory(ctx, story)
	if err != nil {
		return nil, err
	}

	return &CreateStoryRes{
		ID:        res.ID,
		Content:   res.Content,
		Type:      res.ContentType,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}, nil
}

// GetUserStories implements Service.
func (s *service) GetUserStories(ctx context.Context, userID string) ([]*Story, error) {
	panic("unimplemented")
}
