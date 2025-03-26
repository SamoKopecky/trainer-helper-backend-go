package timeslot_handler

import (
	"fmt"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	apiTimeslots, err := cc.TimeslotService.GetByRoleAndDate(params.StartDate, params.EndDate, cc.GetClaims())

	if apiTimeslots == nil {
		apiTimeslots = []model.ApiTimeslot{}
	}

	return cc.JSON(http.StatusOK, apiTimeslots)
}

func Post(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[timeslotPostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	timeslotName := fmt.Sprintf("from %s to %s on %s",
		humanTime(params.Start),
		humanTime(params.End),
		humanDate(params.Start))
	newTimeslot := model.BuildTimeslot(timeslotName, params.Start, params.End, nil, params.TrainerId, nil)
	err = cc.TimeslotCrud.Insert(newTimeslot)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, model.ApiTimeslot{Timeslot: *newTimeslot})
}

func Delete(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[timeslotDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.TimeslotCrud.Delete(params.Id)

	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func Put(c echo.Context) error {
	cc := c.(*schemas.DbContext)

	params, err := api.BindParams[timeslotPutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.TimeslotCrud.Update(&model)
	if err != nil {
		return err
	}
	return cc.NoContent(http.StatusOK)
}
