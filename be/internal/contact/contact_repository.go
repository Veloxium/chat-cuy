package contact

import (
	"context"

	"github.com/Gylmynnn/websocket-sesat/internal/user"
	"github.com/Gylmynnn/websocket-sesat/pkg/utils"
)

type repository struct {
	db utils.DBTX
}

func NewRepository(db utils.DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddContact(ctx context.Context, contact *Contact) (*Contact, error) {
	query := "INSERT INTO contacts (user_id, username, avatar, created_at) VALUES ($1,$2,$3, NOW()) RETURNING id , created_at"
	err := r.db.QueryRowContext(ctx, query, contact.UserId, contact.Username, contact.Avatar).Scan(&contact.ID, &contact.CreatedAt)
	if err != nil {
		return nil, err
	}

	return contact, nil

}

func (r *repository) DeleteContact(ctx context.Context, contactID int64) error {

	query := "UPDATE contacts SET deleted_at = NOW() WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, contactID)
	return err

}

func (r *repository) GetContactWithUser(ctx context.Context, userId int64) (*GetContactsWithUserRes, error) {

	query := `
      SELECT
         u.id, u.username, u.email, u.avatar, u.bio, u.created_at,
         c.id, c.username, c.avatar, c.created_at
      FROM users u
      LEFT JOIN contacts c ON u.id = c.user_id
      WHERE u.id = $1; 
   `

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user user.User
	contacts := []Contact{}

	for rows.Next() {
		var c Contact
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.Avatar, &user.Bio,
			&user.CreatedAt, &c.ID, &c.Username, &c.Avatar, &c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	return &GetContactsWithUserRes{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Contacts:  contacts,
	}, nil

}

func (r *repository) GetContactByUserId(ctx context.Context, userID int64) ([]Contact, error) {

	query := "SELECT id, user_id, username, created_at FROM contacts WHERE user_id = $1 AND deleted_at IS NULL"

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		err := rows.Scan(&c.ID, &c.UserId, &c.Username, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil

}
