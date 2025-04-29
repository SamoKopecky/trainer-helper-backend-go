package work_set

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func PostUndelete(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[workSetUndeletePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	// TODO: Make batch undelete
	for _, id := range params.Ids {
		err = cc.WorkSetCrud.Undelete(id)
	}
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
