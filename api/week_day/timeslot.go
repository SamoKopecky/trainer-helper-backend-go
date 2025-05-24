package weekday

import (
	"net/http"
	"strconv"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func DeleteTimeslot(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.WeekDayCrud.DeleteTimeslot(id)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
