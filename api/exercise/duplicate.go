package exercise

import (
	"maps"
	"net/http"
	"slices"
	"trainer-helper/api"
	"trainer-helper/model"

	"github.com/labstack/echo/v4"
)

// FIXME: Rework this
func PostDuplicate(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseDuplicatePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	// TODO: Service
	newExercises, err := updateExericses(&params, cc)
	if err != nil {
		return err
	}

	if len(newExercises) == 0 {
		newExercises = []*model.Exercise{}
	}

	return cc.JSON(http.StatusOK, newExercises)
}

// TODO: Make service
func updateExericses(param *exerciseDuplicatePostParams, cc *api.DbContext) (newExercises []*model.Exercise, err error) {
	err = cc.ExerciseCrud.DeleteByWeekDayId(param.TimeslotId)
	if err != nil {
		return
	}

	copyExercises, err := cc.ExerciseCrud.GetExerciseWorkSets([]int{param.CopyTimeslotId})
	if err != nil {
		return
	}

	newExercisesMap := make(map[int]*model.Exercise)
	var newWorkSets []model.WorkSet
	for _, exercise := range copyExercises {
		// Adjust new exercise
		exercise.ToNew(param.TimeslotId)

		// NOTE: Possible performance improvment, insert many
		err = cc.ExerciseCrud.Insert(&exercise)
		if err != nil {
			return
		}
		for _, ws := range exercise.WorkSets {
			ws.ToNew(exercise.Id)
			newWorkSets = append(newWorkSets, ws)
		}
		// Clean up
		exercise.WorkSets = nil
		newExercisesMap[exercise.Id] = &exercise
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
