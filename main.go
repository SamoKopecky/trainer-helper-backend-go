package main

import (
	"flag"
	"trainer-helper/api/app"
	"trainer-helper/config"
)

func main() {
	cfg := config.GetConfig()
	seed := flag.Bool("seed", false, "Seed database")
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	dbConn := GetDbConn(cfg, *debug)
	dbConn.RunMigrations()
	if *seed {
		dbConn.SeedDb()
	}

	app.RunApi(dbConn.Conn, &cfg)
}
