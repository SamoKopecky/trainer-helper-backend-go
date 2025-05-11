package api

import (
	"time"
	"trainer-helper/utils"

	"github.com/labstack/echo/v4"
)

func BindParams[T any](c echo.Context) (T, error) {
	params := *new(T)
	if err := c.Bind(&params); err != nil {
		return params, err
	}
	return params, nil
}

func DerefString(ptr *string) string {
	if ptr == nil {
		return ""
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

func DerefDate(ptr *utils.Date) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return (*ptr).Time
}
