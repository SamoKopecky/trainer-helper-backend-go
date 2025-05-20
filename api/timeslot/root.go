package timeslot

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schema"

	"github.com/labstack/echo/v4"
)

func GetManyEnhanced(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotEnchancedGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	users, err := cc.UserService.GetUsers(cc.Claims)
	if err != nil {
		return err
	}

	apiTimeslots, err := cc.TimeslotService.GetByRoleAndDate(
		params.Start,
		params.End,
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

func GetMany(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	timeslots, err := cc.TimeslotService.GetByStartAndUser(params.StartDate, params.EndDate, params.UserId)
	if err != nil {
		return err
	}

	if timeslots == nil {
		timeslots = []model.Timeslot{}
	}

	return cc.JSON(http.StatusOK, timeslots)

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
