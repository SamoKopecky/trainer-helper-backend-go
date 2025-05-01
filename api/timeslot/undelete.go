package timeslot

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostUndelete(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timestlotUndeletePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.TimeslotCrud.UndeleteMany([]int{params.Id})
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
