package work_set_handler

import (
	"trainer-helper/api"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[workSetPutRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	crud := crud.NewCRUDWorkSet(cc.Db)
	model := params.toModel()
	crud.Update(&model)
	return err
}
