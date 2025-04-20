package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DbConn struct {
	Conn      *bun.DB
	Migration *migrate.Migrate
	Dsn       string
}

func GetDbConn(dsn string, debug bool, migrationPath string) DbConn {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}
	return DbConn{
		Conn:      db,
		Dsn:       dsn,
		Migration: m,
	}
}

func (d DbConn) RunMigrations() {
	err := d.Migration.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No changes to apply, continuing...")
	} else if err == nil {
		fmt.Println("Applying migrations...")
	} else {
		log.Fatal(err)
	}
}

func (d DbConn) DownMigrations() {
	err := d.Migration.Down()

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No changes to apply, continuing...")
	} else if err == nil {
		fmt.Println("Applying migrations...")
	} else {
		log.Fatal(err)
	}
}
