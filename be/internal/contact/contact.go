package contact

import (
	"context"
	"time"
)

type Contact struct {
	ID        int64      `json:"id" db:"id"`
	UserId    int64      `json:"user_id" db:"user_id"`
	Username  string     `json:"username" db:"username"`
	Avatar    string     `json:"avatar" db:"avatar"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type CreateContactReq struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Avatar   string `json:"avatar" db:"avatar"`
}

type CreateContactRes struct {
	ID        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Avatar    string    `json:"avatar" db:"avatar"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type GetContactsRes struct {
	ID        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Avatar    string    `json:"avatar" db:"avatar"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type GetContactsWithUserRes struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Bio       string    `json:"bio" db:"bio"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Contacts  []Contact `json:"contacts" db:"contacts"`
}

type Repository interface {
	AddContact(c context.Context, contact *Contact) (*Contact, error)
	DeleteContact(c context.Context, contactID int64) error
	GetAllContacts(c context.Context) ([]GetContactsRes, error)
	GetContactByID(c context.Context, contactID int64) error
	GetContactWithUser(c context.Context, userId int64) (*GetContactsWithUserRes, error)
	GetContactByUserId(c context.Context, userID int64) ([]Contact, error)
}

type Service interface {
	AddContact(c context.Context, req *CreateContactReq) (*CreateContactRes, error)
	DeleteContact(c context.Context, contactID int64) error
	GetAllContacts(c context.Context) ([]GetContactsRes, error)
	GetContactByID(c context.Context, contactID int64) (*GetContactsRes, error)
	// GetContactsByUserID(c context.Context, userID int64) ([]GetContactsRes, error)
	// UpdateContact(c context.Context, contactID int64, req *CreateContactReq) (*CreateContactRes, error)
	// GetContactWithUser(c context.Context, userID int64) (*GetContactsWithUserRes, error)

}
