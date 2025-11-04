package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	authV1API "github.com/HeyReyHR/twitch-clone/iam/internal/api/auth/v1"
	userV1API "github.com/HeyReyHR/twitch-clone/iam/internal/api/user/v1"
	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository"
	authRepository "github.com/HeyReyHR/twitch-clone/iam/internal/repository/auth"
	userRepository "github.com/HeyReyHR/twitch-clone/iam/internal/repository/user"
	"github.com/HeyReyHR/twitch-clone/iam/internal/service"
	authService "github.com/HeyReyHR/twitch-clone/iam/internal/service/auth"
	userService "github.com/HeyReyHR/twitch-clone/iam/internal/service/user"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/closer"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/logger"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/migrator"
	authV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/auth/v1"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

type diContainer struct {
	authV1API authV1.AuthServiceServer
	userV1API userV1.UserServiceServer

	authService service.AuthService
	userService service.UserService

	authRepository repository.AuthRepository
	userRepository repository.UserRepository

	postgresDBConn *pgx.Conn
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) AuthV1API(ctx context.Context) authV1.AuthServiceServer {
	if d.authV1API == nil {
		d.authV1API = authV1API.NewApi(d.AuthService(ctx))
	}

	return d.authV1API
}

func (d *diContainer) UserV1API(ctx context.Context) userV1.UserServiceServer {
	if d.userV1API == nil {
		d.userV1API = userV1API.NewApi(d.UserService(ctx))
	}

	return d.userV1API
}

func (d *diContainer) AuthService(ctx context.Context) service.AuthService {
	if d.authService == nil {
		d.authService = authService.NewService(
			d.AuthRepository(ctx),
			d.UserRepository(ctx),
			config.AppConfig().JWTTokens.AccessTokenTTL(),
			config.AppConfig().JWTTokens.RefreshTokenTTL())
	}

	return d.authService
}

func (d *diContainer) UserService(ctx context.Context) service.UserService {
	if d.userService == nil {
		d.userService = userService.NewService(d.UserRepository(ctx))
	}

	return d.userService
}

func (d *diContainer) AuthRepository(ctx context.Context) repository.AuthRepository {
	if d.authRepository == nil {
		d.authRepository = authRepository.NewRepository(d.PostgresDBConn(ctx))
	}

	return d.authRepository
}

func (d *diContainer) UserRepository(ctx context.Context) repository.UserRepository {
	if d.userRepository == nil {
		d.userRepository = userRepository.NewRepository(d.PostgresDBConn(ctx))
	}

	return d.userRepository
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
