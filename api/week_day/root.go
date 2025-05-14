package weekday

import (
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

	weekDays, err := cc.WeekDayCrud.GetByWeekId(params.WeekId)
	if err != nil {
		return err
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
	return api.DeleteModel(cc, cc.WeekCrud)
}
