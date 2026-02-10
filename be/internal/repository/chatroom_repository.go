package repository

import (
	"context"
	"database/sql"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
)

type ChatRoomRepository interface {
	Create(ctx context.Context, room *model.ChatRoom) error
	GetByID(ctx context.Context, id int) (*model.ChatRoom, error)
	GetByUserID(ctx context.Context, userID int) ([]model.ChatRoomWithLastMessage, error)
	AddMember(ctx context.Context, chatRoomID, userID int) error
	FindPrivateRoom(ctx context.Context, userID1, userID2 int) (*model.ChatRoom, error)
}

type chatRoomRepository struct {
	db *sql.DB
}

func NewChatRoomRepository(db *sql.DB) ChatRoomRepository {
	return &chatRoomRepository{db: db}
}

func (r *chatRoomRepository) Create(ctx context.Context, room *model.ChatRoom) error {
	query := `
		INSERT INTO chat_rooms (name, is_group)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRowContext(ctx, query, room.Name, room.IsGroup).
		Scan(&room.ID, &room.CreatedAt, &room.UpdatedAt)
}

func (r *chatRoomRepository) GetByID(ctx context.Context, id int) (*model.ChatRoom, error) {
	query := `
		SELECT id, name, is_group, created_at, updated_at
		FROM chat_rooms WHERE id = $1`

	room := &model.ChatRoom{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&room.ID, &room.Name, &room.IsGroup, &room.CreatedAt, &room.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *chatRoomRepository) GetByUserID(ctx context.Context, userID int) ([]model.ChatRoomWithLastMessage, error) {
	query := `
		SELECT
			cr.id, cr.name, cr.is_group, cr.created_at, cr.updated_at,
			COALESCE(m.text, '') AS last_message,
			COALESCE(m.created_at, cr.created_at) AS last_message_time,
			COALESCE(m.sender_id, 0) AS last_sender_id,
			u.id AS other_user_id,
			u.username AS other_username,
			COALESCE(u.avatar, '') AS other_avatar,
			COALESCE(u.bio, '') AS other_bio
		FROM chat_rooms cr
		JOIN chat_room_members crm ON cr.id = crm.chat_room_id
		LEFT JOIN LATERAL (
			SELECT text, created_at, sender_id
			FROM messages
			WHERE chat_room_id = cr.id
			ORDER BY created_at DESC
			LIMIT 1
		) m ON true
		LEFT JOIN chat_room_members crm2 ON cr.id = crm2.chat_room_id AND crm2.user_id != $1
		LEFT JOIN users u ON crm2.user_id = u.id
		WHERE crm.user_id = $1
		ORDER BY last_message_time DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.ChatRoomWithLastMessage
	for rows.Next() {
		var room model.ChatRoomWithLastMessage
		if err := rows.Scan(
			&room.ID, &room.Name, &room.IsGroup, &room.CreatedAt, &room.UpdatedAt,
			&room.LastMessage, &room.LastMessageTime, &room.LastSenderID,
			&room.OtherUserID, &room.OtherUsername, &room.OtherAvatar, &room.OtherBio,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, rows.Err()
}

func (r *chatRoomRepository) AddMember(ctx context.Context, chatRoomID, userID int) error {
	query := `
		INSERT INTO chat_room_members (chat_room_id, user_id)
		VALUES ($1, $2)
		ON CONFLICT (chat_room_id, user_id) DO NOTHING`

	_, err := r.db.ExecContext(ctx, query, chatRoomID, userID)
	return err
}

func (r *chatRoomRepository) FindPrivateRoom(ctx context.Context, userID1, userID2 int) (*model.ChatRoom, error) {
	query := `
		SELECT cr.id, cr.name, cr.is_group, cr.created_at, cr.updated_at
		FROM chat_rooms cr
		WHERE cr.is_group = false
		AND cr.id IN (
			SELECT chat_room_id FROM chat_room_members WHERE user_id = $1
		)
		AND cr.id IN (
			SELECT chat_room_id FROM chat_room_members WHERE user_id = $2
		)
		LIMIT 1`

	room := &model.ChatRoom{}
	err := r.db.QueryRowContext(ctx, query, userID1, userID2).Scan(
		&room.ID, &room.Name, &room.IsGroup, &room.CreatedAt, &room.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return room, nil
}
