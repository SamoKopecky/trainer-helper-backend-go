package exercise_type

import (
	"net/http"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func PostDuplicate(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	types, err := cc.ExerciseTypeService.DuplicateDefault(cc.Claims.Subject)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, types)
}
