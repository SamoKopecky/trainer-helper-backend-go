package timeslot

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/schema"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
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

	params, err := api.BindParams[timeslotPostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newTimeslot := params.ToModel()
	err = cc.TimeslotCrud.Insert(&newTimeslot)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, schema.Timeslot{Timeslot: newTimeslot})
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.TimeslotCrud)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotPutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.ToModel()
	err = cc.TimeslotCrud.Update(&model)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
