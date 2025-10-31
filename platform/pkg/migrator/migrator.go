package migrator

import (
	"database/sql"

	"github.com/HeyReyHR/twitch-clone/platform/pkg/migrator/pg"
)

func NewPgMigrator(db *sql.DB, migrationsDir string) *pg.Migrator {
	return pg.NewMigrator(db, migrationsDir)
}
