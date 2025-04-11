package schemas

import (
	"fmt"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/fetcher"
	"trainer-helper/service"
	"trainer-helper/store"

	"github.com/labstack/echo/v4"
)

type DbContext struct {
	echo.Context

	ExerciseCrud     store.Exercise
	TimeslotCrud     store.Timeslot
	WorkSetCrud      store.WorkSet
	ExerciseTypeCrud store.ExerciseType

	IAMFetcher fetcher.IAM

	TimeslotService     service.Timeslot
	PersonService       service.Person
	ExerciseTypeService service.ExerciseType

	Claims *api.JwtClaims
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
