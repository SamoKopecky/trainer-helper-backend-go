package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type DbContext struct {
	Db *bun.DB
	echo.Context
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
