package user

import (
	"github.com/jackc/pgx/v5"

	repository2 "github.com/HeyReyHR/twitch-clone/iam/internal/repository"
)

var _ repository2.UserRepository = (*repository)(nil)

type repository struct {
	dbConn *pgx.Conn
}

func NewRepository(dbConn *pgx.Conn) *repository {
	return &repository{
		dbConn: dbConn,
	}
}
