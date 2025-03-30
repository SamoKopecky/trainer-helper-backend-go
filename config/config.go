package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabasePort     string `env:"DB_PORT" envDefault:"5432"`
	DatabaseHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DatabasePassword string `env:"DB_PASSWORD" envDefault:"alpharius"`
	DatabaseUser     string `env:"DB_USER" envDefault:"trainer_helper"`
	DatabaseName     string `env:"DB_NAME" envDefault:"trainer_helper"`

	Env string `env:"ENV" envDefault:"dev"`

	KeycloakBaseUrl           string `env:"KC_BASE_URL" envDefault:"http://localhost:8080"`
	KeycloakAdminClientId     string `env:"KC_ADMIN_CLIENT_ID" envDefault:"admin-cli"`
	KeycloakAdminClientSecret string `env:"KC_ADMIN_CLIENT_SECRET"`
	KeycloakRealm             string `env:"KC_REALM" envDefault:"trainer-helper"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName)
}

func GetConfig() (c Config) {
	err := env.ParseWithOptions(&c, env.Options{
		Prefix: "APP_",
	})
	if err != nil {
		log.Fatal(err)
	}
	return
}
