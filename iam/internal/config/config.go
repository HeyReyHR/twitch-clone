package config

import (
	"os"

	"github.com/HeyReyHR/twitch-clone/iam/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger    LoggerConfig
	IamGRPC   ServiceConfig
	Postgres  PostgresConfig
	Redis     RedisConfig
	JWTTokens JWTTokensConfig
	Password  PasswordConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	iamCfg, err := env.NewIamGRPCConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	jwtTokensCfg, err := env.NewJWTTokensConfig()
	if err != nil {
		return err
	}

	passwordCfg, err := env.NewPasswordConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:    loggerCfg,
		IamGRPC:   iamCfg,
		Postgres:  postgresCfg,
		Redis:     redisCfg,
		JWTTokens: jwtTokensCfg,
		Password:  passwordCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
