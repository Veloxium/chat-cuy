package model

import "time"

type Contact struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ContactID int       `json:"contact_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactWithUser struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ContactID int       `json:"contact_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type AddContactRequest struct {
	ContactUsername string `json:"contact_username"`
}
