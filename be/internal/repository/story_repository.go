package repository

import (
	"context"
	"database/sql"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
)

type StoryRepository interface {
	Create(ctx context.Context, story *model.Story) error
	GetAll(ctx context.Context) ([]model.StoryWithUser, error)
	GetByUserID(ctx context.Context, userID int) ([]model.StoryWithUser, error)
	Delete(ctx context.Context, id int, userID int) error
}

type storyRepository struct {
	db *sql.DB
}

func NewStoryRepository(db *sql.DB) StoryRepository {
	return &storyRepository{db: db}
}

func (r *storyRepository) Create(ctx context.Context, story *model.Story) error {
	query := `
		INSERT INTO stories (user_id, image_url, caption)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	return r.db.QueryRowContext(ctx, query,
		story.UserID, story.ImageURL, story.Caption,
	).Scan(&story.ID, &story.CreatedAt)
}

func (r *storyRepository) GetAll(ctx context.Context) ([]model.StoryWithUser, error) {
	query := `
		SELECT s.id, s.user_id, u.username, u.avatar, s.image_url, s.caption, s.created_at
		FROM stories s
		JOIN users u ON s.user_id = u.id
		ORDER BY s.created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []model.StoryWithUser
	for rows.Next() {
		var story model.StoryWithUser
		if err := rows.Scan(
			&story.ID, &story.UserID, &story.Username, &story.Avatar,
			&story.ImageURL, &story.Caption, &story.CreatedAt,
		); err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}
	return stories, rows.Err()
}

func (r *storyRepository) GetByUserID(ctx context.Context, userID int) ([]model.StoryWithUser, error) {
	query := `
		SELECT s.id, s.user_id, u.username, u.avatar, s.image_url, s.caption, s.created_at
		FROM stories s
		JOIN users u ON s.user_id = u.id
		WHERE s.user_id = $1
		ORDER BY s.created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []model.StoryWithUser
	for rows.Next() {
		var story model.StoryWithUser
		if err := rows.Scan(
			&story.ID, &story.UserID, &story.Username, &story.Avatar,
			&story.ImageURL, &story.Caption, &story.CreatedAt,
		); err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}
	return stories, rows.Err()
}

func (r *storyRepository) Delete(ctx context.Context, id int, userID int) error {
	query := `DELETE FROM stories WHERE id = $1 AND user_id = $2`
	result, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
