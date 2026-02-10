package service

import (
	"context"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
)

type StoryService interface {
	Create(ctx context.Context, story *model.Story) error
	GetAll(ctx context.Context) ([]model.StoryWithUser, error)
	GetByUserID(ctx context.Context, userID int) ([]model.StoryWithUser, error)
	Delete(ctx context.Context, id int, userID int) error
}

type storyService struct {
	storyRepo repository.StoryRepository
}

func NewStoryService(storyRepo repository.StoryRepository) StoryService {
	return &storyService{storyRepo: storyRepo}
}

func (s *storyService) Create(ctx context.Context, story *model.Story) error {
	return s.storyRepo.Create(ctx, story)
}

func (s *storyService) GetAll(ctx context.Context) ([]model.StoryWithUser, error) {
	return s.storyRepo.GetAll(ctx)
}

func (s *storyService) GetByUserID(ctx context.Context, userID int) ([]model.StoryWithUser, error) {
	return s.storyRepo.GetByUserID(ctx, userID)
}

func (s *storyService) Delete(ctx context.Context, id int, userID int) error {
	return s.storyRepo.Delete(ctx, id, userID)
}
