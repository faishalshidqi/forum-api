package bootstrap

import (
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

type Env struct {
	ServerAddr      string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout  int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenKey  string `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey string `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenAge  int    `mapstructure:"ACCESS_TOKEN_AGE"`
	RefreshTokenAge int    `mapstructure:"REFRESH_TOKEN_AGE"`
	PGUser          string `mapstructure:"PGUSER"`
	PGPassword      string `mapstructure:"PGPASSWORD"`
	PGHost          string `mapstructure:"PGHOST"`
	PGPort          string `mapstructure:"PGPORT"`
	PGDatabase      string `mapstructure:"PGDATABASE"`

	MaxConnections    int
	MinConnections    int
	MaxConnLifeTime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Error reading config file", slog.String("error", err.Error()))
		return nil
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		slog.Error("unable to decode into struct", slog.String("error", err.Error()))
		return nil
	}
	return &env
}
