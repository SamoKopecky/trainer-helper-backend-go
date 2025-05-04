package week

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newWeek := params.toModel(cc.Claims.Subject)
	err = cc.WeekCrud.Insert(newWeek)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newWeek)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.WeekCrud.Update(params.toModel())
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.WeekCrud)
}
