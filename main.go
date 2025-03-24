package main

import (
	"flag"
	"log"
	"trainer-helper/api/app"

	"github.com/caarlos0/env/v11"
)

func main() {
	seed := flag.Bool("seed", false, "Seed database")
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	var cfg Config
	err := env.ParseWithOptions(&cfg, env.Options{
		Prefix: "APP_",
	})
	if err != nil {
		log.Fatal(err)
	}

	dbConn := GetDbConn(cfg, *debug)
	dbConn.RunMigrations()
	if *seed {
		dbConn.SeedDb()
	}

	app.RunApi(dbConn.Conn)
}
