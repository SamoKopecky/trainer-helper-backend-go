package weekday

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostFromRaw(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekDayPostFromRawRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	trainerId := cc.Claims.Subject
	exercises, err := cc.AIService.GenerateWeekDay(trainerId, params.RawData, params.WeekDayId)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, exercises)
}
