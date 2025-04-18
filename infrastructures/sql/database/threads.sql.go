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

const getThreadById = `-- name: GetThreadById :one
select threads.id, threads.title, threads.body, threads.date, users.username from threads join users on threads.owner = users.id where threads.id = $1
`

type GetThreadByIdRow struct {
	ID       pgtype.UUID      `db:"id" json:"id"`
	Title    string           `db:"title" json:"title"`
	Body     string           `db:"body" json:"body"`
	Date     pgtype.Timestamp `db:"date" json:"date"`
	Username string           `db:"username" json:"username"`
}

func (q *Queries) GetThreadById(ctx context.Context, id pgtype.UUID) (GetThreadByIdRow, error) {
	row := q.db.QueryRow(ctx, getThreadById, id)
	var i GetThreadByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Date,
		&i.Username,
	)
	return i, err
}
