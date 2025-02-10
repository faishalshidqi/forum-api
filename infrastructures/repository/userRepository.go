package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
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
			Name:     user.Fullname,
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

func (ur *userRepository) Fetch(ctx context.Context) ([]domains.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (domains.User, error) {
	user, err := ur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return domains.User{}, err
	}
	return domains.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Password:  user.Password,
		FullName:  user.Name,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (domains.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(database bootstrap.Database) domains.UserRepository {
	return &userRepository{
		database: database,
	}
}
