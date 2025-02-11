package domains

import (
	"context"
	"forum-api/commons/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository interface {
	Add(ctx context.Context, user SignupRequest) (SignupResponseData, error)
	Fetch(ctx context.Context) ([]database.User, error)
	GetByUsername(ctx context.Context, username string) (database.User, error)
	GetByID(ctx context.Context, id pgtype.UUID) (database.User, error)
}
