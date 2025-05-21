package weekday

import (
	"errors"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekDayGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	var weekDays []model.WeekDay

	if params.WeekId != nil {
		weekDays, err = cc.WeekDayCrud.GetByWeekIdsWithDeleted([]int{*params.WeekId})
		if err != nil {
			return err
		}
	} else if params.DayDate != nil && params.UserId != nil {
		weekDays, err = cc.WeekDayCrud.GetByDate(*&params.DayDate.Time, *params.UserId)
		if err != nil {
			return err
		}
	} else {
		return cc.BadRequest(errors.New("Provide either 'week_id' OR both 'day_date' AND 'user_id'"))
	}

	if weekDays == nil {
		return cc.JSON(http.StatusOK, []model.Week{})
	}

	return cc.JSON(http.StatusOK, weekDays)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[weekDayPostRequest](cc, cc.WeekDayCrud)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PutModel[weekDayPutRequest](cc, cc.WeekDayCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.WeekDayCrud)
}
