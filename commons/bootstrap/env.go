package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerAddr      string `mapstructure:"SERVER_ADDRESS"`
	PGUser          string `mapstructure:"PGUSER"`
	PGPassword      string `mapstructure:"PGPASSWORD"`
	PGHost          string `mapstructure:"PGHOST"`
	PGPort          string `mapstructure:"PGPORT"`
	PGDatabase      string `mapstructure:"PGDATABASE"`
	AccessTokenKey  string `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey string `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenAge  int    `mapstructure:"ACCESS_TOKEN_AGE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &env
}
