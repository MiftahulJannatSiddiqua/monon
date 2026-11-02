package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDb(db *sqlx.DB, dir string) error {
	migrates := &migrate.FileMigrationSource{
		Dir: dir,
	}
	_, err := migrate.Exec(db.DB, "postgres", migrates, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("Successfully migrated database")
	return nil
}