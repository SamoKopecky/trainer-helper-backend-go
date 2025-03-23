package exercise_handler

import (
	"log"
	"maps"
	"net/http"
	"slices"
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

	res, err := cc.CRUDExercise.GetExerciseWorkSets(params.Id)
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice of points so that we can append worksets
	exercisesMap := make(map[int32]*model.ExerciseWorkSets)

	for _, r := range res {
		val, ok := exercisesMap[r.ExerciseId]
		if !ok {
			val = &model.ExerciseWorkSets{
				Exercise: r.ToExercise(),
			}
			exercisesMap[r.ExerciseId] = val
		}
		val.WorkSets = append(val.WorkSets, r.ToWorkSet())
		val.WorkSetCount += 1

	}
	exercises := slices.Collect(maps.Values(exercisesMap))
	if len(exercises) == 0 {
		exercises = []*model.ExerciseWorkSets{}
	}
	apiTimeslot, err := cc.CRUDTimeslot.GetById(params.Id)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: sort

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
	newWorkSets := make([]model.WorkSet, 2)
	for i := range 2 {
		newWorkSets[i] = *model.BuildWorkSet(newExercise.Id, 0, (*int32)(nil), "-")
	}
	err = cc.CRUDWorkSet.InsertMany(&newWorkSets)
	if err != nil {
		log.Fatal(err)
	}

	return cc.JSON(http.StatusOK, model.ExerciseWorkSets{
		Exercise:     *newExercise,
		WorkSetCount: int32(len(newWorkSets)),
		WorkSets:     newWorkSets,
	})

}
