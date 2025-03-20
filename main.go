package main

import (
	"flag"
	"trainer-helper/api"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	seed := flag.Bool("seed", false, "Seed database")
	flag.Parse()

	dbConn := GetDbConn()
	dbConn.RunMigrations()
	if *seed {
		dbConn.SeedDb()
	}

	api.RunApi(dbConn.Conn)
}
