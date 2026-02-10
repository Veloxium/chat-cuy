package model

import "time"

type Story struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

type StoryWithUser struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}
