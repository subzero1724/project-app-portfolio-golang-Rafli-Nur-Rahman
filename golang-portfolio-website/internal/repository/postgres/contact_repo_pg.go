package postgres

import (
	"context"
	"portfolio/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ContactRepositoryPG struct {
	db *pgxpool.Pool
}

func NewContactRepository(db *pgxpool.Pool) *ContactRepositoryPG {
	return &ContactRepositoryPG{db: db}
}

func (r *ContactRepositoryPG) Create(ctx context.Context, contact *model.Contact) error {
	query := `
		INSERT INTO contacts (name, email, subject, message, is_read)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	return r.db.QueryRow(ctx, query,
		contact.Name,
		contact.Email,
		contact.Subject,
		contact.Message,
		false,
	).Scan(&contact.ID, &contact.CreatedAt)
}

func (r *ContactRepositoryPG) GetByID(ctx context.Context, id string) (*model.Contact, error) {
	query := `
		SELECT id, name, email, subject, message, is_read, created_at
		FROM contacts WHERE id = $1
	`
	contact := &model.Contact{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&contact.ID,
		&contact.Name,
		&contact.Email,
		&contact.Subject,
		&contact.Message,
		&contact.IsRead,
		&contact.CreatedAt,
	)
	return contact, err
}

func (r *ContactRepositoryPG) GetAll(ctx context.Context) ([]*model.Contact, error) {
	query := `
		SELECT id, name, email, subject, message, is_read, created_at
		FROM contacts ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*model.Contact
	for rows.Next() {
		contact := &model.Contact{}
		err := rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.Email,
			&contact.Subject,
			&contact.Message,
			&contact.IsRead,
			&contact.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactRepositoryPG) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE contacts SET is_read = $1 WHERE id = $2`
	isRead := status == "read"
	_, err := r.db.Exec(ctx, query, isRead, id)
	return err
}

func (r *ContactRepositoryPG) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM contacts WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
