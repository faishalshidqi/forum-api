package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type postgresTaskRepository struct {
	database bootstrap.Database
}

func (ptr *postgresTaskRepository) Add(c context.Context, task domains.Task) (domains.AddTaskResponseData, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(task.Owner)
	if err != nil {
		return domains.AddTaskResponseData{}, err
	}
	thread, err := ptr.database.Query.CreateThread(c, database.CreateThreadParams{
		Title: task.Title,
		Body:  task.Body,
		Owner: uuid,
	})
	if err != nil {
		return domains.AddTaskResponseData{}, err
	}
	return thread.ToAddTaskResponseData(), nil
}

func NewPostgresTaskRepository(database bootstrap.Database) domains.TaskRepository {
	return &postgresTaskRepository{
		database: database,
	}
}
