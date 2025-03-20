package main

import (
	"flag"
	"trainer-helper/api"
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

	api.RunApi(dbConn.Conn)
}
