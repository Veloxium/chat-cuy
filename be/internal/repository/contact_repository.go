package repository

import (
	"context"
	"database/sql"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
)

type ContactRepository interface {
	Add(ctx context.Context, userID, contactID int) error
	GetByUserID(ctx context.Context, userID int) ([]model.ContactWithUser, error)
	Exists(ctx context.Context, userID, contactID int) (bool, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Add(ctx context.Context, userID, contactID int) error {
	query := `
		INSERT INTO contacts (user_id, contact_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, contact_id) DO NOTHING`

	_, err := r.db.ExecContext(ctx, query, userID, contactID)
	return err
}

func (r *contactRepository) GetByUserID(ctx context.Context, userID int) ([]model.ContactWithUser, error) {
	query := `
		SELECT c.id, c.user_id, c.contact_id,
			u.username, u.email, COALESCE(u.bio, '') AS bio,
			COALESCE(u.avatar, '') AS avatar, c.created_at
		FROM contacts c
		JOIN users u ON c.contact_id = u.id
		WHERE c.user_id = $1
		ORDER BY u.username ASC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []model.ContactWithUser
	for rows.Next() {
		var contact model.ContactWithUser
		if err := rows.Scan(
			&contact.ID, &contact.UserID, &contact.ContactID,
			&contact.Username, &contact.Email, &contact.Bio,
			&contact.Avatar, &contact.CreatedAt,
		); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, rows.Err()
}

func (r *contactRepository) Exists(ctx context.Context, userID, contactID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM contacts WHERE user_id = $1 AND contact_id = $2)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, userID, contactID).Scan(&exists)
	return exists, err
}
