package repository

import (
	"context"
	"database/sql"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	SearchByUsername(ctx context.Context, query string) ([]model.User, error)
	UpdateAvatar(ctx context.Context, userID int, avatarURL string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (username, email, password, bio, avatar)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		user.Username, user.Email, user.Password, user.Bio, user.Avatar,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, username, email, password, bio, avatar, created_at, updated_at
		FROM users WHERE email = $1`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.Bio, &user.Avatar, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `
		SELECT id, username, email, password, bio, avatar, created_at, updated_at
		FROM users WHERE id = $1`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.Bio, &user.Avatar, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
		SELECT id, username, email, password, bio, avatar, created_at, updated_at
		FROM users WHERE username = $1`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.Bio, &user.Avatar, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) SearchByUsername(ctx context.Context, query string) ([]model.User, error) {
	sqlQuery := `
		SELECT id, username, email, bio, avatar, created_at, updated_at
		FROM users WHERE username ILIKE $1
		LIMIT 20`

	rows, err := r.db.QueryContext(ctx, sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID, &user.Username, &user.Email,
			&user.Bio, &user.Avatar, &user.CreatedAt, &user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (r *userRepository) UpdateAvatar(ctx context.Context, userID int, avatarURL string) error {
	query := `UPDATE users SET avatar = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, avatarURL, userID)
	return err
}
