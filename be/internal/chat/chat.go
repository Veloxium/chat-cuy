package chat

import (
	"context"
	"time"
)

type Chat struct {
	ID        string    `json:"id" db:"id"`
	ChatName  string    `json:"chat_name" db:"chat_name"`
	IsGroup   bool      `json:"is_group" db:"is_group"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateChatReq struct {
	ChatName string `json:"chat_name"`
	IsGroup  bool   `json:"is_group"`
}

type CreateChatRes struct {
	ChatName  string    `json:"chat_name"`
	IsGroup   bool      `json:"is_group"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateMessageReq struct {
	ChatID   string `json:"chat_id"`
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
}

type CreateMessageRes struct {
	ID          int       `json:"id"`
	ChatID      string    `json:"chat_id"`
	SenderID    string    `json:"sender_id"`
	Content     string    `json:"content"`
	MessageType string    `json:"message_type"`
	MediaURL    *string   `json:"media_url,omitempty"`
	IsRead      bool      `json:"is_read"`
	IsReply     bool      `json:"is_reply"`
	CreatedAt   time.Time `json:"created_at"`
}

type Message struct {
	ID          int        `json:"id" db:"id"`
	ChatID      string     `json:"chat_id" db:"chat_id"`
	SenderID    string     `json:"sender_id" db:"sender_id"`
	Content     string     `json:"content" db:"content"`
	MessageType string     `json:"message_type" db:"message_type"`
	MediaURL    *string    `json:"media_url,omitempty" db:"media_url"`
	IsRead      bool       `json:"is_read" db:"is_read"`
	IsReply     bool       `json:"is_reply" db:"is_reply"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Repository interface {
	CreateChat(ctx context.Context, chat *Chat) (*Chat, error)
	CreateMessage(ctx context.Context, message *Message) (*Message, error)
	GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error)
}

type Service interface {
	CreateChat(ctx context.Context, chat *CreateChatReq) (*CreateChatRes, error)
	CreateMessage(ctx context.Context, message *CreateMessageReq) (*CreateMessageRes, error)
	GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error)
}
