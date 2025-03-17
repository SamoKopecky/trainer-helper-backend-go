package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"trainer-helper/crud"
	"trainer-helper/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DbContext struct {
	Db *bun.DB
	echo.Context
}

var db *bun.DB

func main() {
	dsn := "postgres://root:alpharius@localhost/trainer_helper?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	runMigrations(sqldb, dsn)
	// seedDb(*db)
	crud := crud.CRUDTimeslot{Db: db}
	timeslots, err := crud.GetByTimeRange(time.Now().Add(-2*24*time.Hour), time.Now().Add(10*24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res :", len(timeslots))
	api()
}

func api() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &DbContext{Db: db, Context: c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.GET("/-/ping", func(c echo.Context) error {
		cc := c.(*DbContext)
		crud := crud.CRUDTimeslot{Db: db}
		timeslots, err := crud.GetByTimeRange(time.Now().Add(-2*24*time.Hour), time.Now().Add(10*24*time.Hour))
		if err != nil {
			log.Fatal(err)
		}
		return cc.JSON(http.StatusOK, timeslots)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func runMigrations(sql *sql.DB, dsn string) {
	driver, _ := postgres.WithInstance(sql, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dsn, driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}

func seedDb(db bun.DB) {
	ctx := context.Background()
	const TRAINER_ID = 1
	var timeslots []model.Timeslot
	timeNow := time.Now()

	for i := 0; i < 7; i++ {
		timeslots = append(timeslots, *model.BuildTimeslot("some name", timeNow, timeNow.Add(1*time.Hour), TRAINER_ID, nil))
		timeNow = timeNow.Add(24 * time.Hour)
	}

	res, err := db.NewInsert().Model(&timeslots).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)

}
