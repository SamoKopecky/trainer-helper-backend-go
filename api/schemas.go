package api

import (
	"fmt"
	"net/http"
	"trainer-helper/crud"
	"trainer-helper/service"

	"github.com/labstack/echo/v4"
)

type DbContext struct {
	echo.Context

	CRUDExercise crud.CRUDExercise
	CRUDTimeslot crud.CRUDTimeslot
	CRUDWorkSet  crud.CRUDWorkSet

	IAM service.IAM
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
