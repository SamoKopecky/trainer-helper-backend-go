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

	blocks, err := cc.WeekService.GetBlocks(params.userId)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, blocks)
}
