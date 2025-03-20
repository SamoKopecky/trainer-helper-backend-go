package api

import (
	"log"
	"net/http"
	"time"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunApi(db *bun.DB) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &dbContext{Db: db, Context: c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.GET("/-/ping", test)
	e.Logger.Fatal(e.Start(":1323"))
}

func test(c echo.Context) error {
	cc := c.(*dbContext)
	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslots, err := crud.GetByTimeRange(time.Now().Add(-2*24*time.Hour), time.Now().Add(10*24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, timeslots)
}
