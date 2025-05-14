package exercise

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func GetMany(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[exerciseGetParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	exercises, err := cc.ExerciseService.GetExerciseWorkSets(params.WeekDayIds)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, exercises)
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PutModel[exercisePutParams](cc, cc.ExerciseCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.ExerciseCrud)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exercisePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	// Create exercise
	newExercise := model.BuildExercise(params.TimeslotId, params.GroupId, nil, nil)
	err = cc.ExerciseCrud.Insert(newExercise)
	if err != nil {
		return err
	}

	// Create worksets
	const workSetCount = 1
	newWorkSets := make([]model.WorkSet, workSetCount)
	for i := range workSetCount {
		newWorkSets[i] = *model.BuildWorkSet(newExercise.Id, 0, (*int)(nil), "-")
	}
	err = cc.WorkSetCrud.InsertMany(&newWorkSets)
	if err != nil {
		return err
	}

	newExercise.WorkSets = newWorkSets
	return cc.JSON(http.StatusOK, newExercise)
}
