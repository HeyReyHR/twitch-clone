package app

import (
	"context"
	"fmt"

	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/cache"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/cache/redis"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/closer"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/logger"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/migrator"
	authV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/auth/v1"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

type diContainer struct {
	authV1API authV1.AuthServiceServer
	userV1API userV1.UserServiceServer

	postgresDBConn *pgx.Conn

	redisPool   *redigo.Pool
	redisClient cache.RedisClient
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) PostgresDBConn(ctx context.Context) *pgx.Conn {
	if d.postgresDBConn == nil {
		dbConn, dbErr := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if dbErr != nil {
			panic(fmt.Sprintf("❌ failed to connect to Postgres: %s\n", dbErr.Error()))
		}

		dbErr = dbConn.Ping(ctx)
		if dbErr != nil {
			panic(fmt.Sprintf("❌ failed ping database: %s\n", dbErr.Error()))
		}

		migratorRunner := migrator.NewPgMigrator(stdlib.OpenDB(*dbConn.Config().Copy()), config.AppConfig().Postgres.MigrationsDir())
		dbErr = migratorRunner.Up()
		if dbErr != nil {
			logger.Error(ctx, "❌ failed to run migrations", zap.Error(dbErr))
			return nil
		}

		closer.AddNamed("Postgres conn", func(ctx context.Context) error {
			return dbConn.Close(ctx)
		})

		d.postgresDBConn = dbConn
	}

	return d.postgresDBConn
}

func (d *diContainer) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}

	return d.redisPool
}

func (d *diContainer) RedisClient() cache.RedisClient {
	if d.redisClient == nil {
		d.redisClient = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}

	return d.redisClient
}
