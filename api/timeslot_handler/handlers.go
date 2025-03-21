package timeslot_handler

import (
	"fmt"
	"log"
	"net/http"
	"trainer-helper/api"
	"trainer-helper/crud"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslots, err := crud.GetByTimeRange(params.StartDate, params.EndDate)
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, timeslots)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotPostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	crudTimeslot := crud.CRUDTimeslot{Db: cc.Db}
	crudPerson := crud.CRUDPerson{Db: cc.Db}

	timeslotName := fmt.Sprintf("from %s to %s on %s",
		humanTime(params.StartDate),
		humanTime(params.EndDate),
		humanDate(params.StartDate))
	newTimeslot := model.BuildTimeslot(timeslotName, params.StartDate, params.EndDate, params.TrainerId, nil)
	err = crudTimeslot.Insert(newTimeslot)
	if err != nil {
		log.Fatal(err)
	}

	return cc.JSON(http.StatusOK, toFullTimeslot(newTimeslot, crudPerson))
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	crudTimeslot := crud.CRUDTimeslot{Db: cc.Db}
	crudPerson := crud.CRUDPerson{Db: cc.Db}
	timeslot, err := crudTimeslot.Delete(params.Id)

	if err != nil {
		log.Fatal(err)
	}

	if timeslot.IsEmpty() {
		return cc.NoContent(http.StatusNotFound)
	}

	return cc.JSON(http.StatusOK, toFullTimeslot(timeslot, crudPerson))
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[timeslotPutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	model := params.toModel()
	crud.Update(&model)
	return err
}
