package exercise_type_duplicate_handler

import (
	"net/http"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	types, err := cc.ExerciseTypeService.DuplicateDefault(cc.Claims.Subject)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, types)
}
