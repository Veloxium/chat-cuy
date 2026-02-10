package repository

import (
	"context"
	"database/sql"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *model.Message) error
	GetByChatRoomID(ctx context.Context, chatRoomID int) ([]model.Message, error)
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(ctx context.Context, msg *model.Message) error {
	query := `
		INSERT INTO messages (text, sender_id, chat_room_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query,
		msg.Text, msg.SenderID, msg.ChatRoomID,
	).Scan(&msg.ID, &msg.CreatedAt, &msg.UpdatedAt)
}

func (r *messageRepository) GetByChatRoomID(ctx context.Context, chatRoomID int) ([]model.Message, error) {
	query := `
		SELECT id, text, sender_id, chat_room_id, created_at, updated_at
		FROM messages
		WHERE chat_room_id = $1
		ORDER BY created_at ASC`

	rows, err := r.db.QueryContext(ctx, query, chatRoomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(
			&msg.ID, &msg.Text, &msg.SenderID, &msg.ChatRoomID,
			&msg.CreatedAt, &msg.UpdatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, rows.Err()
}
