package bootstrap

import (
	"context"
	"fmt"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Database struct {
	Pool  *pgxpool.Pool
	Query *database.Queries
}

func NewPSQLDatabase(env *Env) *Database {
	psqlUser := env.PGUser
	psqlPassword := env.PGPassword
	psqlHost := env.PGHost
	psqlPort := env.PGPort
	psqlDatabase := env.PGDatabase

	env.MaxConnections = 10
	env.MinConnections = 0
	env.MaxConnLifeTime = 30 * time.Minute
	env.MaxConnIdleTime = 10 * time.Minute
	env.HealthCheckPeriod = 2 * time.Minute

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", psqlUser, psqlPassword, psqlHost, psqlPort, psqlDatabase)
	//db, err := sql.Open("postgres", dbUrl)
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Pool:  db,
		Query: database.New(db),
	}
}
