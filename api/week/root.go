package week

import (
	"database/sql"
	"errors"
	"net/http"
	"trainer-helper/api"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[WeekGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	week, err := cc.WeekCrud.GetClosestToDate(params.StartDate.Time, params.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return cc.JSON(http.StatusOK, map[string]string{})
	}

	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, week)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[weekPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newWeek := params.ToModel()
	err = cc.WeekService.CreateWeek(&newWeek, params.IsFirst)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newWeek)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PutModel[weekPutRequest](cc, cc.WeekCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.WeekCrud)
}
