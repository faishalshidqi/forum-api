package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/commons/sql/database"
	"forum-api/domains"
	"github.com/jackc/pgx/v5/pgtype"
)

type userRepository struct {
	database bootstrap.Database
}

func (ur *userRepository) Add(ctx context.Context, user domains.SignupRequest) (domains.SignupResponseData, error) {
	returned, err := ur.database.Query.CreateUser(
		ctx,
		database.CreateUserParams{
			Username: user.Username,
			Password: user.Password,
			Fullname: user.Fullname,
		},
	)
	if err != nil {
		return domains.SignupResponseData{}, err
	}
	returnedData := domains.SignupResponseData{}
	returnedData.ID = returned.ID.String()
	returnedData.Username = user.Username
	returnedData.FullName = user.Fullname
	return returnedData, nil
}

func (ur *userRepository) Fetch(ctx context.Context) ([]database.User, error) {
	users, err := ur.database.Query.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (database.User, error) {
	user, err := ur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id pgtype.UUID) (database.User, error) {
	user, err := ur.database.Query.GetByID(ctx, id)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}

func NewUserRepository(database bootstrap.Database) domains.UserRepository {
	return &userRepository{
		database: database,
	}
}
