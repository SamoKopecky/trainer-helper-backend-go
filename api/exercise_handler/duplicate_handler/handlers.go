package exercise_duplicate_handler

import (
	"maps"
	"net/http"
	"slices"
	"time"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseDuplicatePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	newTimeslot, err := updateTimeslot(params, cc)
	if err != nil {
		return err
	}
	newExercises, err := updateExericses(params, cc)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, model.TimeslotExercises{
		Exercises: newExercises,
		Timeslot:  newTimeslot,
	})
}

func updateExericses(param *exerciseDuplicatePostParams, cc *api.DbContext) (newExercises []*model.Exercise, err error) {
	err = cc.CRUDExercise.DeleteByTimeslot(param.TimeslotId)
	if err != nil {
		return
	}

	copyExercises, err := cc.CRUDExercise.GetExerciseWorkSetsTwo(param.CopyTimeslotId)
	if err != nil {
		return
	}

	newExercisesMap := make(map[int32]*model.Exercise)
	var newWorkSets []*model.WorkSet
	for _, e := range copyExercises {
		// Adjust new exercise
		// NOTE: Possible performance improvment, insert many
		e.Id = model.EmptyId
		e.TimeslotId = param.TimeslotId
		e.CreatedAt = time.Now()
		e.UpdatedAt = time.Now()

		err = cc.CRUDExercise.Insert(e)
		if err != nil {
			return
		}
		for _, ws := range e.WorkSets {
			ws.Id = model.EmptyId
			ws.ExerciseId = e.Id
			ws.CreatedAt = time.Now()
			ws.UpdatedAt = time.Now()
			newWorkSets = append(newWorkSets, ws)
		}
		// Clean up
		e.WorkSets = nil
		newExercisesMap[e.Id] = e
	}

	err = cc.CRUDWorkSet.InsertMany(&newWorkSets)
	if err != nil {
		return
	}

	for _, ws := range newWorkSets {
		if e, ok := newExercisesMap[ws.ExerciseId]; ok {
			e.WorkSets = append(e.WorkSets, ws)
		}
	}
	newExercises = slices.Collect(maps.Values(newExercisesMap))
	return
}

func updateTimeslot(params *exerciseDuplicatePostParams, cc *api.DbContext) (newApiTimeslot model.ApiTimeslot, err error) {
	copyTimeslot, err := cc.CRUDTimeslot.GetById(params.CopyTimeslotId)
	if err != nil {
		return
	}

	newTimeslot := model.Timeslot{
		IdModel: model.IdModel{
			Id: params.TimeslotId,
		},
		Name: copyTimeslot.Name,
	}

	err = cc.CRUDTimeslot.Update(&newTimeslot)
	if err != nil {
		return
	}

	newApiTimeslot, err = cc.CRUDTimeslot.GetById(params.TimeslotId)
	return
}
