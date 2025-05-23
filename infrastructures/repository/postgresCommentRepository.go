package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type postgresCommentRepository struct {
	database bootstrap.Database
}

func (pcr *postgresCommentRepository) GetById(c context.Context, commentId string) (domains.Comment, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(commentId)
	if err != nil {
		return domains.Comment{}, err
	}
	comment, err := pcr.database.Query.GetCommentById(c, uuid)
	if err != nil {
		return domains.Comment{}, err
	}
	return comment.ToDomainsComment(), nil
}

func (pcr *postgresCommentRepository) SoftDelete(c context.Context, id string) error {
	commentId := pgtype.UUID{}
	err := commentId.Scan(id)
	if err != nil {
		return err
	}
	return pcr.database.Query.SoftDeleteComment(c, commentId)
}

func (pcr *postgresCommentRepository) GetByThread(c context.Context, thread string) ([]domains.GetCommentsByThreadResponseData, error) {
	threadId := pgtype.UUID{}
	err := threadId.Scan(thread)
	if err != nil {
		return nil, err
	}
	comments, err := pcr.database.Query.GetCommentsByThread(c, threadId)
	if err != nil {
		return nil, err
	}
	returned := make([]domains.GetCommentsByThreadResponseData, 0)
	for _, comment := range comments {
		returned = append(returned, comment.ToGetThreadComments())
	}
	return returned, nil
}

func (pcr *postgresCommentRepository) Add(c context.Context, commentRequest domains.AddCommentRequest, owner, thread string) (domains.AddCommentResponseData, error) {
	ownerId := pgtype.UUID{}
	err := ownerId.Scan(owner)
	if err != nil {
		return domains.AddCommentResponseData{}, err
	}
	threadId := pgtype.UUID{}
	err = threadId.Scan(thread)
	if err != nil {
		return domains.AddCommentResponseData{}, err
	}
	comment, err := pcr.database.Query.CreateComment(c, database.CreateCommentParams{
		Owner:   ownerId,
		Thread:  threadId,
		Content: commentRequest.Content,
	})
	if err != nil {
		return domains.AddCommentResponseData{}, err
	}
	return comment.ToAddCommentResponseData(), nil
}

func NewPostgresCommentRepository(database bootstrap.Database) domains.CommentRepository {
	return &postgresCommentRepository{
		database: database,
	}
}
