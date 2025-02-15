// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: mark_feed_fetched.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = $2, updated_at = $3 WHERE id = $1
`

type MarkFeedFetchedParams struct {
	ID            uuid.UUID
	LastFetchedAt sql.NullTime
	UpdatedAt     time.Time
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.ID, arg.LastFetchedAt, arg.UpdatedAt)
	return err
}
