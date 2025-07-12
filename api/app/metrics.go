package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func RunMetricsApi() {
	e := echo.New()
	e.GET("/-/metrics", echoprometheus.NewHandler())

	if err := e.Start(":5000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
