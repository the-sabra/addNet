package outbox

import (
	"time"

	"github.com/google/uuid"
)

type Outbox struct {
	ID        uuid.UUID  `db:"id"`
	SentAt    *time.Time `db:"sent_at"`
	CreatedAt time.Time  `db:"created_at"`
	Value     int32      `db:"value"`
}
