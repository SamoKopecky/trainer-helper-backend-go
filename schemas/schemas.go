package schemas

import (
	"fmt"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DbContext struct {
	echo.Context

	ExerciseCrud crud.Exercise
	TimeslotCrud crud.Timeslot
	WorkSetCrud  crud.WorkSet

	IAMFetcher fetcher.IAM

	TimeslotService service.Timeslot
	PersonService   service.Person
}

func (c DbContext) GetClaims() *api.JwtClaims {
	return c.Get("user").(*jwt.Token).Claims.(*api.JwtClaims)
}

func (c DbContext) BadRequest(err error) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": fmt.Sprint(err)})
}
