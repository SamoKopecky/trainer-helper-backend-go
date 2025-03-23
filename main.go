package main

import (
	"flag"
	"trainer-helper/api/app"
)

func main() {
	seed := flag.Bool("seed", false, "Seed database")
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	dbConn := GetDbConn(*debug)
	dbConn.RunMigrations()
	if *seed {
		dbConn.SeedDb()
	}

	app.RunApi(dbConn.Conn)
}
