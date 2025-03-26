package exercise_handler

import (
	"net/http"
	"sort"
	"strconv"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	paramId, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}
	params := exerciseGetParams{
		Id: int32(paramId),
	}

	exercises, err := cc.ExerciseCrud.GetExerciseWorkSets(params.Id)
	if err != nil {
		return err
	}
	if exercises == nil {
		exercises = []*model.Exercise{}
	}

	sort.Slice(exercises, func(i, j int) bool {
		if exercises[i].GroupId == exercises[j].GroupId {
			return exercises[i].Id < exercises[j].Id
		}
		return exercises[i].GroupId < exercises[j].GroupId
	})
	for _, exercise := range exercises {
		exercise.SortWorkSets()
	}

	apiTimeslot, err := cc.TimeslotService.GetById(params.Id)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, model.TimeslotExercises{
		Timeslot:  apiTimeslot,
		Exercises: exercises,
	})

}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exercisePutParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.toModel()
	err = cc.ExerciseCrud.Update(&model)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.ExerciseCrud.DeleteByExerciseAndTimeslot(params.TimeslotId, params.ExerciseId)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exercisePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	// Create exercise
	newExercise := model.BuildExercise(params.TimeslotId, params.GroupId, "", model.None)
	err = cc.ExerciseCrud.Insert(newExercise)
	if err != nil {
		return err
	}

	// Create worksets
	const workSetCount = 2
	newWorkSets := make([]*model.WorkSet, workSetCount)
	for i := range workSetCount {
		newWorkSets[i] = model.BuildWorkSet(newExercise.Id, 0, (*int32)(nil), "-")
	}
	err = cc.WorkSetCrud.InsertMany(&newWorkSets)
	if err != nil {
		return err
	}

	newExercise.WorkSets = newWorkSets
	return cc.JSON(http.StatusOK, newExercise)
}
