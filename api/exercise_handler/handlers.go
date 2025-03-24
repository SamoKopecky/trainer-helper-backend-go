package exercise_handler

import (
	"log"
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

	// Get timeslot
	apiTimeslot, err := cc.CRUDTimeslot.GetById(params.Id)
	if err != nil {
		log.Fatal(err)
	}

	exercises, err := cc.CRUDExercise.GetExerciseWorkSetsTwo(params.Id)
	if err != nil {
		log.Fatal(err)
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
	err = cc.CRUDExercise.Update(&model)
	if err != nil {
		log.Fatal(err)
	}

	return cc.NoContent(http.StatusOK)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseDeleteParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.CRUDExercise.DeleteByExerciseAndTimeslot(params.TimeslotId, params.ExerciseId)
	if err != nil {
		log.Fatal(err)
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
	err = cc.CRUDExercise.Insert(newExercise)
	if err != nil {
		log.Fatal(err)
	}

	// Create worksets
	const workSetCount = 2
	newWorkSets := make([]*model.WorkSet, workSetCount)
	for i := range workSetCount {
		newWorkSets[i] = model.BuildWorkSet(newExercise.Id, 0, (*int32)(nil), "-")
	}
	err = cc.CRUDWorkSet.InsertMany(&newWorkSets)
	if err != nil {
		log.Fatal(err)
	}

	newExercise.WorkSets = newWorkSets
	return cc.JSON(http.StatusOK, newExercise)
}
