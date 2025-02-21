// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_tokens.sql

package database

import (
	"context"
)

const addToken = `-- name: AddToken :one
insert into refresh_tokens (token) values($1) returning token
`

func (q *Queries) AddToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRow(ctx, addToken, token)
	err := row.Scan(&token)
	return token, err
}

const deleteToken = `-- name: DeleteToken :one
delete from refresh_tokens where token = $1 returning token
`

func (q *Queries) DeleteToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRow(ctx, deleteToken, token)
	err := row.Scan(&token)
	return token, err
}

const getToken = `-- name: GetToken :one
select token from refresh_tokens where token = $1
`

func (q *Queries) GetToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRow(ctx, getToken, token)
	err := row.Scan(&token)
	return token, err
}
