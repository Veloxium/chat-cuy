package user

import (
	"context"
	"github.com/Gylmynnn/websocket-sesat/utils"
)

type repository struct {
	db utils.DBTX
}

func NewRepository(db utils.DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Putser(ctx context.Context, userID string, req *User) (*User, error) {
	query := `
		UPDATE users
		SET
			username = COALESCE(NULLIF($1, ''), username),
			profile_picture = COALESCE(NULLIF($2, ''), profile_picture),
			about_message = COALESCE(NULLIF($3, ''), about_message),
			updated_at = NOW()
		WHERE id = $4
		RETURNING id, username, profile_picture, about_message, updated_at
	`
	row := r.db.QueryRowContext(ctx, query,
		req.Username,
		req.ProfilePicture,
		req.AboutMessage,
		userID,
	)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.ProfilePicture,
		&user.AboutMessage,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	var user User
	query := `SELECT 
    id,
    email,
    username,
    profile_picture,
    about_message,
    created_at 
    FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.ProfilePicture,
		&user.AboutMessage,
		&user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) SearchUsersByUsername(ctx context.Context, usernamePrefix string) ([]*User, error) {
	var users []*User
	query := `SELECT 
    id,
    username,
    email,
    profile_picture,
    about_message,
    created_at 
    FROM users WHERE username ILIKE $1`
	rows, err := r.db.QueryContext(ctx, query, "%"+usernamePrefix+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.ProfilePicture,
			&user.AboutMessage,
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId string
	query := `INSERT INTO users(
    username,
    email,
    password,
    profile_picture,
    about_message,
    created_at)
    VALUES ($1, $2, $3, $4, $5, NOW()) returning id, created_at`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.ProfilePicture, user.AboutMessage).Scan(&lastInsertId, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.ID = lastInsertId
	return user, nil

}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT 
    id,
    email,
    username,
    password,
    profile_picture,
    about_message,
    created_at 
    FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.ProfilePicture,
		&user.AboutMessage,
		&user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
