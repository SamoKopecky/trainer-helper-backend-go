package work_set_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[workSetPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.WorkSetCrud.Update(&model)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
