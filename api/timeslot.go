package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"trainer-helper/crud"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

type timeslotGetParams struct {
	StartDate time.Time `query:"start_date"`
	EndDate   time.Time `query:"end_date"`
}

type timeslotPostParams struct {
	TrainerId int32     `json:"trainer_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type timeslotDeleteParams struct {
	TimeslotId int32 `json:"timeslot_id"`
}

func timeslotGet(c echo.Context) error {
	cc := c.(*dbContext)

	params, err := bindParams[timeslotGetParams](cc)
	if err != nil {
		return cc.badRequest(err)
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslots, err := crud.GetByTimeRange(params.StartDate, params.EndDate)
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, timeslots)
}

func timeslotPost(c echo.Context) error {
	cc := c.(*dbContext)

	params, err := bindParams[timeslotPostParams](cc)
	if err != nil {
		return cc.badRequest(err)
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslotName := fmt.Sprintf("from %s to %s on %s",
		humanTime(params.StartDate),
		humanTime(params.EndDate),
		humanDate(params.StartDate))
	newTimeslot := model.BuildTimeslot(timeslotName, params.StartDate, params.EndDate, params.TrainerId, nil)
	err = crud.Insert(newTimeslot)
	if err != nil {
		log.Fatal(err)
	}

	return cc.JSON(http.StatusOK, newTimeslot)
}

func timeslotDelete(c echo.Context) error {
	cc := c.(*dbContext)

	params, err := bindParams[timeslotDeleteParams](cc)
	if err != nil {
		return cc.badRequest(err)
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslot, err := crud.Delete(params.TimeslotId)

	if err != nil {
		log.Fatal(err)
	}

	if timeslot.IsEmpty() {
		return cc.NoContent(http.StatusNotFound)
	}

	return cc.JSON(http.StatusOK, timeslot)
}
