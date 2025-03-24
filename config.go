package main

import "fmt"

type Config struct {
	DatabasePort     string `env:"DB_PORT" envDefault:"5432"`
	DatabaseHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DatabasePassword string `env:"DB_PASSWORD" envDefault:"alpharius"`
	DatabaseUser     string `env:"DB_USER" envDefault:"root"`
	DatabaseName     string `env:"DB_NAME" envDefault:"trainer_helper"`
	Env              string `env:"ENV" envDefault:"dev"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName)
}
