package main

import (
	"trainer-helper/api"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// TODO: Add argument parser for seeding
	dbConn := GetDbConn()
	dbConn.RunMigrations()
	// dbConn.SeedDb()

	api.RunApi(dbConn.Conn)
}
