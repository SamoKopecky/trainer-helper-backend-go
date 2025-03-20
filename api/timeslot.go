package api

import (
	"log"
	"net/http"
	"time"
	"trainer-helper/crud"

	"github.com/labstack/echo/v4"
)

type getParams struct {
	StartDate time.Time `query:"start_date"`
	EndDate   time.Time `query:"end_date"`
}

func timeslotGet(c echo.Context) error {
	cc := c.(*dbContext)
	params := new(getParams)

	if err := cc.Bind(params); err != nil {
		return cc.badRequest()
	}

	crud := crud.CRUDTimeslot{Db: cc.Db}
	timeslots, err := crud.GetByTimeRange(params.StartDate, params.EndDate)
	if err != nil {
		log.Fatal(err)
	}
	return cc.JSON(http.StatusOK, timeslots)
}
