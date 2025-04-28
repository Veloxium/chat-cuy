package story

import (
	"context"
	"time"
)

type Story struct {
	ID          int        `db:"id"`
	UserID      string     `db:"user_id"`
	ContentType string     `db:"content_type"`
	Content     string     `db:"content"`
	CreatedAt   time.Time  `db:"created_at"`
	ExpiresAt   time.Time  `db:"expires_at"`
	DeletedAt   *time.Time `db:"deleted_at,omitempty"`
}

type CreateStoryReq struct {
	UserID      string `form:"user_id" binding:"required,uuid"`
	ContentType string `form:"content_type" binding:"required,oneof=text image video"`
	Content     string `form:"content" binding:"required"`
}

type CreateStoryRes struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Type      string    `json:"content_type"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Repository interface {
	CreateStory(ctx context.Context, story *Story) (*Story, error)
	GetActiveStoriesByUserID(ctx context.Context, userID string) ([]*Story, error)
}

type Service interface {
	CreateStory(ctx context.Context, req *CreateStoryReq) (*CreateStoryRes, error)
	GetUserStories(ctx context.Context, userID string) ([]*Story, error)
}
