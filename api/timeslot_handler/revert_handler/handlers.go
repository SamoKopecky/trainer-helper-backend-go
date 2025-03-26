package timeslot_revert_handler

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timestlotRevertPutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.TimeslotCrud.RevertSolfDelete(params.Id)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
