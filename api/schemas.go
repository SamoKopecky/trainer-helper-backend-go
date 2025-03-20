package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type dbContext struct {
	Db *bun.DB
	echo.Context
}
