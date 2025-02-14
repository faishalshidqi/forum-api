package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type postgresThreadRepository struct {
	database bootstrap.Database
}

func (ptr *postgresThreadRepository) GetById(c context.Context, id string) (domains.Thread, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(id)
	if err != nil {
		return domains.Thread{}, err
	}
	thread, err := ptr.database.Query.GetThreadById(c, uuid)
	if err != nil {
		return domains.Thread{}, err
	}
	return thread.ToDomainsThread(), nil
}

func (ptr *postgresThreadRepository) Add(c context.Context, task domains.AddThreadRequest, owner string) (domains.AddThreadResponseData, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(owner)
	if err != nil {
		return domains.AddThreadResponseData{}, err
	}
	thread, err := ptr.database.Query.CreateThread(c, database.CreateThreadParams{
		Title: task.Title,
		Body:  task.Body,
		Owner: uuid,
	})
	if err != nil {
		return domains.AddThreadResponseData{}, err
	}
	return thread.ToAddThreadResponseData(), nil
}

func NewPostgresThreadRepository(database bootstrap.Database) domains.ThreadRepository {
	return &postgresThreadRepository{
		database: database,
	}
}
