package model

import "time"

type Message struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	SenderID   int       `json:"senderId"`
	ChatRoomID int       `json:"chatRoomId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type SendMessageRequest struct {
	RoomID  int    `json:"roomId"`
	Content string `json:"content"`
	UserID  int    `json:"userId"`
}
