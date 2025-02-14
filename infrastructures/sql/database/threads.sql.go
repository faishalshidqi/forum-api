// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: threads.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createThread = `-- name: CreateThread :one
insert into threads (id, title, body, date, owner) values(gen_random_uuid(), $1, $2, now(), $3) returning id, title, owner
`

type CreateThreadParams struct {
	Title string      `db:"title" json:"title"`
	Body  string      `db:"body" json:"body"`
	Owner pgtype.UUID `db:"owner" json:"owner"`
}

type CreateThreadRow struct {
	ID    pgtype.UUID `db:"id" json:"id"`
	Title string      `db:"title" json:"title"`
	Owner pgtype.UUID `db:"owner" json:"owner"`
}

func (q *Queries) CreateThread(ctx context.Context, arg CreateThreadParams) (CreateThreadRow, error) {
	row := q.db.QueryRow(ctx, createThread, arg.Title, arg.Body, arg.Owner)
	var i CreateThreadRow
	err := row.Scan(&i.ID, &i.Title, &i.Owner)
	return i, err
}
