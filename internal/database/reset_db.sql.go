// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reset_db.sql

package database

import (
	"context"
)

const resetDB = `-- name: ResetDB :exec
DELETE FROM users
`

func (q *Queries) ResetDB(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetDB)
	return err
}
