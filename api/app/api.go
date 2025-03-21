package app

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/api/timeslot_handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func RunApi(db *bun.DB) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &api.DbContext{Db: db, Context: c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/-/ping", pong)
	e.GET("/timeslot", timeslot_handler.Get)
	e.POST("/timeslot", timeslot_handler.Post)
	e.DELETE("/timeslot", timeslot_handler.Delete)
	e.PUT("/timeslot", timeslot_handler.Put)

	e.Logger.Fatal(e.Start(":1323"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
