package model

import "time"

type ChatRoom struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsGroup     bool      `json:"is_group"`
	UnreadCount int       `json:"unreadCount"`
	Messages    []Message `json:"messages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ChatRoomMember struct {
	ID         int       `json:"id"`
	ChatRoomID int       `json:"chat_room_id"`
	UserID     int       `json:"user_id"`
	JoinedAt   time.Time `json:"joined_at"`
}

type ChatRoomWithLastMessage struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	IsGroup         bool      `json:"is_group"`
	UnreadCount     int       `json:"unreadCount"`
	LastMessage     string    `json:"lastMessage"`
	LastMessageTime time.Time `json:"lastMessageTime"`
	LastSenderID    int       `json:"lastSenderId"`
	OtherUserID     int       `json:"otherUserId"`
	OtherUsername   string    `json:"otherUsername"`
	OtherAvatar     string    `json:"otherAvatar"`
	OtherBio        string    `json:"otherBio"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
