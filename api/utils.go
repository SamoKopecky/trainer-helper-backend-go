package api

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func bindParams[T any](c echo.Context) (*T, error) {
	params := new(T)
	if err := c.Bind(params); err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	return params, nil
}

func humanTime(time time.Time) string {
	return time.Format("15:04")
}

func humanDate(time time.Time) string {
	return time.Format("02-01")
}
