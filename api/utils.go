package api

import (
	"time"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func BindParams[T any](c echo.Context) (*T, error) {
	params := new(T)
	if err := c.Bind(params); err != nil {
		return nil, err
	}
	return params, nil
}

func DerefString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func DerefExerciseType(ptr *model.ExerciseTypeEnum) model.ExerciseTypeEnum {
	if ptr == nil {
		return model.None
	}
	return *ptr
}

func DerefInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr

}

func DerefTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}
