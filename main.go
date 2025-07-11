package main

import (
	"flag"
	"log"
	"trainer-helper/api/app"
	"trainer-helper/config"
	"trainer-helper/db"

	"github.com/joho/godotenv"
)

const migrationPath = "file://migrations"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v. Continuing without .env", err)
	}
	cfg := config.GetConfig()
	debug := flag.Bool("debug", false, "Show database queries")
	flag.Parse()

	dbConn := db.GetDbConn(cfg.GetDSN(), *debug, migrationPath)
	dbConn.RunMigrations()

	app.RunApi(dbConn.Conn, &cfg)
}
