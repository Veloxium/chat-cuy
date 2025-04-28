package user

import (
	"context"
	"time"
)

type User struct {
	ID             string     `db:"id"`
	Username       string     `db:"username"`
	Email          string     `db:"email"`
	Password       string     `db:"password"`
	ProfilePicture string     `db:"profile_picture"`
	AboutMessage   string     `db:"about_message"`
	IsOnline       bool       `db:"is_online"`
	LastSeen       time.Time  `db:"last_seen"`
	CreatedAt      time.Time  `db:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
}

type SearchUsersRes struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	AboutMessage   string `json:"about_message"`
}

type FindUserByIDRes struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	AboutMessage   string `json:"about_message"`
}

type CreateUserReq struct {
	Username string `form:"username" validate:"required,min=3,max=16"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
	// ProfilePicture string `form:"profile_picture"`
	// AboutMessage   string `form:"about_message" validate:"max=255"`
}

type CreateUserRes struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	AboutMessage   string    `json:"about_message"`
	CreatedAt      time.Time `json:"created_at"`
}

type LoginUserReq struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

type LoginUserRes struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	AboutMessage   string    `json:"about_message"`
	CreatedAt      time.Time `json:"created_at"`
	AccessToken    string    `json:"accessToken"`
}

type LoginUserWithGoogleReq struct {
	AccessToken string `form:"accessToken"`
}

type LoginUserWithGoogleRes struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	AboutMessage   string    `json:"about_message"`
	CreatedAt      time.Time `json:"created_at"`
	AccessToken    string    `json:"accessToken"`
}

type LoginUserWithFacebookReq struct {
	AccessToken string `form:"accessToken"`
}

type UpdateUserReq struct {
	Username       *string `form:"username" validate:"omitempty,min=3,max=16"`
	ProfilePicture *string `form:"-"`
	AboutMessage   *string `form:"about_message" validate:"omitempty,max=255"`
}

type UpdateUserRes struct {
	ID        string     `json:"id"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LoginUserWithFacebookRes struct {
	AccessToken    string    `json:"accessToken"`
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	AboutMessage   string    `json:"about_message"`
	CreatedAt      time.Time `json:"created_at"`
}

type Repository interface {
	Putser(ctx context.Context, userID string, user *User) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	SearchUsersByUsername(ctx context.Context, usernamePrefix string) ([]*User, error)
}

type Service interface {
	UpdateUser(ctx context.Context, userID string, req *UpdateUserReq) (*UpdateUserRes, error)
	FindUserByID(ctx context.Context, userID string) (*FindUserByIDRes, error)
	SearchUsers(ctx context.Context, usernamePrefix string) ([]*SearchUsersRes, error)
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
	LoginWithGoogle(ctx context.Context, req *LoginUserWithGoogleReq) (*LoginUserWithGoogleRes, error)
	LoginWithFacebook(ctx context.Context, req *LoginUserWithFacebookReq) (*LoginUserWithFacebookRes, error)
}
