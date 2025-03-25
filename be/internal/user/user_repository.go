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
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.ProfilePicture, &user.AboutMessage, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
