package user

import (
	"context"
	"time"
)

type User struct {
	ID             string     `json:"id" db:"id"`
	Username       string     `json:"username" db:"username"`
	Email          string     `json:"email" db:"email"`
	Password       string     `json:"password" db:"password"`
	ProfilePicture string     `json:"profile_picture" db:"profile_picture"`
	AboutMessage   string     `json:"about_message" db:"about_message"`
	IsOnline       bool       `json:"is_online" db:"is_online"`
	LastSeen       time.Time  `json:"last_seen" db:"last_seen"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRes struct {
	ID             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	AboutMessage   string    `json:"about_message" db:"about_message"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserRes struct {
	AccessToken    string    `json:"accessToken"`
	ID             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	AboutMessage   string    `json:"about_message" db:"about_message"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type LoginUserWithGoogleReq struct {
	AccessToken string `json:"accessToken"`
}

type LoginUserWithGoogleRes struct {
	AccessToken    string    `json:"accessToken"`
	ID             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	AboutMessage   string    `json:"about_message" db:"about_message"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type LoginUserWithFacebookReq struct {
	AccessToken string `json:"accessToken"`
}

type LoginUserWithFacebookRes struct {
	AccessToken    string    `json:"accessToken"`
	ID             string    `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	AboutMessage   string    `json:"about_message" db:"about_message"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
	LoginWithGoogle(c context.Context, req *LoginUserWithGoogleReq) (*LoginUserWithGoogleRes, error)
	LoginWithFacebook(c context.Context, req *LoginUserWithFacebookReq) (*LoginUserWithFacebookRes, error)
}
