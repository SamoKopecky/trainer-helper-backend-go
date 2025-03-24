package timeslot_handler

import (
	"fmt"
	"log"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	timeslots, err := cc.CRUDTimeslot.GetByTimeRange(params.StartDate, params.EndDate)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: If empty return empty array
	return cc.JSON(http.StatusOK, timeslots)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotPostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	timeslotName := fmt.Sprintf("from %s to %s on %s",
		humanTime(params.Start),
		humanTime(params.End),
		humanDate(params.Start))
	newTimeslot := model.BuildTimeslot(timeslotName, params.Start, params.End, nil, params.TrainerId, nil)
	err = cc.CRUDTimeslot.Insert(newTimeslot)
	if err != nil {
		log.Fatal(err)
	}

	return cc.JSON(http.StatusOK, toFullTimeslot(newTimeslot, cc.CRUDPerson))
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	timeslot, err := cc.CRUDTimeslot.Delete(params.Id)

	if err != nil {
		log.Fatal(err)
	}

	if timeslot.IsEmpty() {
		return cc.NoContent(http.StatusNotFound)
	}

	return cc.JSON(http.StatusOK, toFullTimeslot(timeslot, cc.CRUDPerson))
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotPutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.CRUDTimeslot.Update(&model)
	if err != nil {
		log.Fatal(err)
	}
	return cc.NoContent(http.StatusOK)
}
