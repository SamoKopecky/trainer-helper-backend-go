package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"trainer-helper/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DbConn struct {
	Conn   *bun.DB
	Driver *sql.DB
	Dsn    string
}

func GetDbConn(debug bool) DbConn {
	// TODO: Make config
	dsn := "postgres://root:alpharius@localhost/trainer_helper?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	if debug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}
	return DbConn{
		Conn:   db,
		Driver: sqldb,
		Dsn:    dsn,
	}
}

func (d DbConn) RunMigrations() {
	driver, err := postgres.WithInstance(d.Driver, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		d.Dsn, driver)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("No changes to apply, continuing...")
	} else if err == nil {
		fmt.Println("Applying migrations...")
	} else {
		log.Fatal(err)
	}

}

func (d DbConn) SeedDb() {
	personId := d.seedUsers()
	d.seedTimeslots(personId)

}

func (d DbConn) seedUsers() int32 {
	ctx := context.Background()

	user := *model.BuildPerson("Samo Kopecky", "abc@wow.com")
	_, err := d.Conn.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		panic(err)
	}
	return user.Id
}

func (d DbConn) seedTimeslots(personId int32) {
	ctx := context.Background()
	const TRAINER_ID = 1
	var timeslots []model.Timeslot
	timeNow := time.Now()

	for range 7 {
		timeslots = append(timeslots, *model.BuildTimeslot("some name", timeNow, timeNow.Add(1*time.Hour), TRAINER_ID, &personId))
		timeNow = timeNow.Add(24 * time.Hour)
	}

	_, err := d.Conn.NewInsert().Model(&timeslots).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
