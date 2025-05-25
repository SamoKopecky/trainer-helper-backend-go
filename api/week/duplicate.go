package week

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostDuplicate(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[WeekDuplicatePostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.WeekService.DuplicateWeekDays(params.TemplateWeekId, params.NewWeekId)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
