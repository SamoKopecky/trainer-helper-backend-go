package weekday

import (
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func PostUndelete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostUndeleteModel(cc, cc.WeekDayCrud)
}
