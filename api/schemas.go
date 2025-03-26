package api

import (
	"fmt"
	"net/http"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/service"

	"github.com/labstack/echo/v4"
)

type DbContext struct {
	echo.Context

	// TODO: Rename crud to just the name without crud
	ExerciseCrud crud.Exercise
	TimeslotCrud crud.Timeslot
	WorkSetCrud  crud.WorkSet

	IAMFetcher fetcher.IAM

	TimeslotService service.Timeslot
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
