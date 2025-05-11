package weekday

import (
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

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
