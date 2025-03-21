package api

import (
	"github.com/labstack/echo/v4"
)

func BindParams[T any](c echo.Context) (*T, error) {
	params := new(T)
	if err := c.Bind(params); err != nil {
		return nil, err
	}
	return params, nil
}
