package exercise

import (
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostUndelete(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[exerciseUndeletePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.ExerciseCrud.UndeleteMany([]int{params.Id})
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
