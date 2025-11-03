package auth

import (
	"github.com/jackc/pgx/v5"
)

type repository struct {
	dbConn *pgx.Conn
}

func NewRepository(dbConn *pgx.Conn) *repository {
	return &repository{
		dbConn: dbConn,
	}
}
