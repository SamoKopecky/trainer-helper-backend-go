package timeslot

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/schema"

	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	users, err := cc.UserService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}

	apiTimeslots, err := cc.TimeslotService.GetByRoleAndDate(
		params.StartDate,
		params.EndDate,
		users,
		cc.Claims)
	if err != nil {
		return err
	}

	if apiTimeslots == nil {
		apiTimeslots = []schema.Timeslot{}
	}

	return cc.JSON(http.StatusOK, apiTimeslots)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[timeslotPostParams](cc, cc.TimeslotCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.TimeslotCrud)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PutModel[timeslotPutParams](cc, cc.TimeslotCrud)
}
