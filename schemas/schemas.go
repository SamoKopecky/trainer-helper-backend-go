package schemas

import (
	"fmt"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/service"

	"github.com/labstack/echo/v4"
)

type DbContext struct {
	echo.Context

	ExerciseCrud crud.Exercise
	TimeslotCrud crud.Timeslot
	WorkSetCrud  crud.WorkSet
	SetTypeCrud  crud.SetType

	IAMFetcher fetcher.IAM

	TimeslotService service.Timeslot
	PersonService   service.Person

	Claims *api.JwtClaims
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
