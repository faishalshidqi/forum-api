// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: comments.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addComment = `-- name: AddComment :one
insert into comments(id, owner, thread, content, date) values(gen_random_uuid(), $1, $2, $3, now()) returning id, content, owner
`

type AddCommentParams struct {
	Owner   pgtype.UUID `db:"owner" json:"owner"`
	Thread  pgtype.UUID `db:"thread" json:"thread"`
	Content string      `db:"content" json:"content"`
}

type AddCommentRow struct {
	ID      pgtype.UUID `db:"id" json:"id"`
	Content string      `db:"content" json:"content"`
	Owner   pgtype.UUID `db:"owner" json:"owner"`
}

func (q *Queries) AddComment(ctx context.Context, arg AddCommentParams) (AddCommentRow, error) {
	row := q.db.QueryRow(ctx, addComment, arg.Owner, arg.Thread, arg.Content)
	var i AddCommentRow
	err := row.Scan(&i.ID, &i.Content, &i.Owner)
	return i, err
}
