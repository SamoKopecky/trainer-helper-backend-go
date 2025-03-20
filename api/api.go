package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type dbContext struct {
	Db *bun.DB
	echo.Context
}

func (c dbContext) badRequest() error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters"})
}

func RunApi(db *bun.DB) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &dbContext{Db: db, Context: c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())

	e.GET("/-/ping", pong)
	e.GET("/timeslot", timeslotGet)

	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
