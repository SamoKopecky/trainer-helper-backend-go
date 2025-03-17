package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"trainer-helper/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://root:alpharius@localhost/trainer_helper?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	ctx := context.Background()

	db := bun.NewDB(sqldb, pgdialect.New())
	var num int
	err := db.QueryRowContext(ctx, "SELECT 1").Scan(&num)
	if err != nil {
		log.Fatal(err)
	}

	println("Result:", num)

	run_migrations(sqldb, dsn)

	person := model.BuildPerson("a", "b")
	res, err := db.NewInsert().Model(person).Exec(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	user := new(model.Person)
	err = db.NewSelect().Model(user).Where("id = ?", 1).Scan(ctx)
	fmt.Printf("%+v\n", user)

}

func run_migrations(sql *sql.DB, dsn string) {
	driver, _ := postgres.WithInstance(sql, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dsn, driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
