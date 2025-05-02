package week

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	blocks, err := cc.WeekService.GetBlocks(params.UserId)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, blocks)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[weekPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newWeek := params.toModel()
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

	params, err := api.BindParams[weekDeleteRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.WeekCrud.Delete(params.Id)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
