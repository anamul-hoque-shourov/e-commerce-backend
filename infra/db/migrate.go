package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migration := &migrate.FileMigrationSource{
		Dir: dir,
	}

	_, err := migrate.Exec(db.DB, "postgres", migration, migrate.Up)
	if err != nil {
		fmt.Println("Migration failed:", err)
		return err
	}
	
	fmt.Println("Migration successful")
	return nil 
}
