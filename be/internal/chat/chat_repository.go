package chat

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

func (r *repository) CreateChat(ctx context.Context, chat *Chat) (*Chat, error) {
	query := `INSERT INTO chats(
	chat_name,
	is_group,
	created_at)
    VALUES ($1, $2, NOW()) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, chat.ChatName, chat.IsGroup).Scan(&chat.ID)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// CreateMessage implements Repository.
func (r *repository) CreateMessage(ctx context.Context, message *Message) (*Message, error) {
	query := `INSERT INTO messages (
    chat_id,
    sender_id,
    content,
    created_at) 
    VALUES ($1, $2, $3, NOW()) RETURNING id, created_at`
	err := r.db.QueryRowContext(
		context.Background(),
		query,
		message.ChatID,
		message.SenderID,
		message.Content).
		Scan(&message.ID, &message.CreatedAt)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *repository) GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error) {
	query := `SELECT 
    id,
    chat_id,
    sender_id,
    content,
    message_type,
    media_url,
    is_read,
    is_reply,
    created_at
    FROM messages WHERE chat_id = $1 ORDER BY created_at ASC`
	rows, err := r.db.QueryContext(ctx, query, chatID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []*Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID,
			&msg.ChatID,
			&msg.SenderID,
			&msg.Content,
			&msg.MessageType,
			&msg.MediaURL,
			&msg.IsRead,
			&msg.IsReply,
			&msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}
	return messages, nil
}
