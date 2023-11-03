package test

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func Migrate(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file:///Users/yamanakajunichi/study/go-unittest-architecture/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	if err = migrate.Up(); err != nil {
		return err
	}

	return nil
}
