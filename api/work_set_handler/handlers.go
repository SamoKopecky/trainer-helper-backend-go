package work_set_handler

import (
	"log"
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[workSetPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.CRUDWorkSet.Update(&model)
	if err != nil {
		log.Fatal(err)
	}
	return cc.NoContent(http.StatusOK)
}
