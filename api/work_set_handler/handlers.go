package work_set_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[[]workSetPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	models := make([]model.WorkSet, len(*params))
	for i, param := range *params {
		models[i] = param.toModel()
	}
	err = cc.WorkSetCrud.UpdateMany(models)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
