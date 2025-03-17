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
	var lastInsertId int
	query := `INSERT INTO users(username, email, password, avatar, bio, created_at)
    VALUES ($1, $2, $3, $4, $5, NOW()) returning id, created_at`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password, user.Avatar, user.Bio).Scan(&lastInsertId, &user.CreatedAt)
	if err != nil {
		return &User{}, err
	}
	user.ID = int64(lastInsertId)
	return user, nil

}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := User{}
	query := "SELECT id, email, username, password, avatar, bio, created_at FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Avatar, &user.Bio, &user.CreatedAt)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}
