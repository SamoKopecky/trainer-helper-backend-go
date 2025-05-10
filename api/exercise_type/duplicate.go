package exercise_type

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostDuplicate(c echo.Context) error {
	cc := c.(*api.DbContext)

	types, err := cc.ExerciseTypeService.DuplicateDefault(cc.Claims.Subject)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, types)
}
