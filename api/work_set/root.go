package work_set

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[[]workSetPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	models := make([]model.WorkSet, len(params))
	for i, param := range params {
		models[i] = param.ToModel()
	}
	err = cc.WorkSetCrud.UpdateMany(models)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
