package main

import (
	"flag"
	"trainer-helper/api/app"
	"trainer-helper/config"
	"trainer-helper/db"
)

const migrationPath = "file://migrations"

func main() {
	cfg := config.GetConfig()
	seed := flag.Bool("seed", false, "Seed database")
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	dbConn := db.GetDbConn(cfg.GetDSN(), *debug, migrationPath)
	dbConn.RunMigrations()
	if *seed {
		dbConn.SeedDb()
	}

	app.RunApi(dbConn.Conn, &cfg)
}
