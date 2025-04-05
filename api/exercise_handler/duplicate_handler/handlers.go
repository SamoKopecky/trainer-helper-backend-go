package exercise_duplicate_handler

import (
	"maps"
	"net/http"
	"slices"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schemas"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*schemas.DbContext)
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

	if len(newExercises) == 0 {
		newExercises = []*model.Exercise{}
	}

	return cc.JSON(http.StatusOK, model.TimeslotExercises{
		Exercises: newExercises,
		Timeslot:  newTimeslot,
	})
}

func updateExericses(param *exerciseDuplicatePostParams, cc *schemas.DbContext) (newExercises []*model.Exercise, err error) {
	err = cc.ExerciseCrud.DeleteByTimeslot(param.TimeslotId)
	if err != nil {
		return
	}

	copyExercises, err := cc.ExerciseCrud.GetExerciseWorkSets(param.CopyTimeslotId)
	if err != nil {
		return
	}

	newExercisesMap := make(map[int]*model.Exercise)
	var newWorkSets []model.WorkSet
	for _, exercise := range copyExercises {
		// Adjust new exercise
		exercise.ToNew(param.TimeslotId)

		// NOTE: Possible performance improvment, insert many
		err = cc.ExerciseCrud.Insert(exercise)
		if err != nil {
			return
		}
		for _, ws := range exercise.WorkSets {
			ws.ToNew(exercise.Id)
			newWorkSets = append(newWorkSets, ws)
		}
		// Clean up
		exercise.WorkSets = nil
		newExercisesMap[exercise.Id] = exercise
	}

	err = cc.WorkSetCrud.InsertMany(&newWorkSets)
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

func updateTimeslot(params *exerciseDuplicatePostParams, cc *schemas.DbContext) (newApiTimeslot model.ApiTimeslot, err error) {
	copyTimeslot, err := cc.TimeslotCrud.GetById(params.CopyTimeslotId)
	if err != nil {
		return
	}

	newTimeslot := model.Timeslot{
		IdModel: model.IdModel{
			Id: params.TimeslotId,
		},
		Name: copyTimeslot.Name,
	}

	err = cc.TimeslotCrud.Update(&newTimeslot)
	if err != nil {
		return
	}

	newApiTimeslot, err = cc.TimeslotService.GetById(params.TimeslotId)
	return
}
