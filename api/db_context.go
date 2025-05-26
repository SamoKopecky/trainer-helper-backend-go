package api

import (
	"fmt"
	"net/http"
	"trainer-helper/config"
	"trainer-helper/fetcher"
	"trainer-helper/schema"
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
	BlockCrud        store.Block
	WeekCrud         store.Week
	WeekDayCrud      store.WeekDay

	IAMFetcher fetcher.IAM

	TimeslotService     service.Timeslot
	UserService         service.User
	ExerciseTypeService service.ExerciseType
	BlockService        service.Block
	WeekService         service.Week
	ExerciseService     service.Exercise

	Claims *schema.JwtClaims

	Config *config.Config
}

func (c DbContext) BadRequest(err error) error {
	errStr := fmt.Sprint(err)
	// TODO: log error too
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": errStr})
}
