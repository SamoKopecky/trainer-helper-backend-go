package exercise_count_handler

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseCountPostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newWorkSets := make([]*model.WorkSet, params.Count)
	for i := range params.Count {
		newWorkSet := params.WorkSetTemplate
		newWorkSet.Id = model.EmptyId
		newWorkSets[i] = &newWorkSet
	}
	err = cc.CRUDWorkSet.InsertMany(&newWorkSets)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, newWorkSets)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseCountDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	ids, err := cc.CRUDWorkSet.DeleteMany(params.WorkSetIds)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, ids)
}
