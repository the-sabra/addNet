package outbox

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Store(ctx context.Context, value int32) error
	GetPendingMessages(ctx context.Context, limit int) ([]Outbox, error)
	MarkAsProcessed(ctx context.Context, id uuid.UUID) error
}

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) Store(ctx context.Context, value int32) error {
	query := `INSERT INTO outbox (id,value) VALUES ($1, $2)`
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, query, id, value)
	if err != nil {
		log.Println("Error inserting into outbox:", err)
		return err
	}
	return nil
}

func (r *SQLRepository) GetPendingMessages(ctx context.Context, limit int) ([]Outbox, error) {
	query := `SELECT id, sent_at, created_at, value FROM outbox WHERE sent_at IS NULL LIMIT $1`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outboxes []Outbox
	for rows.Next() {
		var outbox Outbox
		if err := rows.Scan(&outbox.ID, &outbox.SentAt, &outbox.CreatedAt, &outbox.Value); err != nil {
			return nil, err
		}
		outboxes = append(outboxes, outbox)
	}

	return outboxes, nil
}

func (r *SQLRepository) MarkAsProcessed(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE outbox SET sent_at = $1 WHERE id = $2`
	sentAt := time.Now()
	_, err := r.db.ExecContext(ctx, query, sentAt, id)
	if err != nil {
		return err
	}
	return nil
}
