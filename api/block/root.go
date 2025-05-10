package block

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[blockGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	blocks, err := cc.BlockService.GetBlocks(params.UserId)
	if err != nil {
		return err
	}

	if blocks == nil {
		return cc.JSON(http.StatusOK, []model.Block{})
	}

	return cc.JSON(http.StatusOK, blocks)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[blockPostRequest](cc, cc.BlockCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.BlockCrud)
}
