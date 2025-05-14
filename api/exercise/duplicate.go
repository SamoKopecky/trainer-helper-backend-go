package exercise

import (
	"net/http"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/schema"

	"github.com/labstack/echo/v4"
)

func PostDuplicate(c echo.Context) error {
	cc := c.(*api.DbContext)
	params, err := api.BindParams[exerciseDuplicatePostParams](cc)
	if err != nil {
		return cc.BadRequest(err)
	}
	// TODO: Service
	// newTimeslot, err := updateTimeslot(&params, cc)
	// if err != nil {
	// 	return err
	// }
	// TODO: Service
	newExercises, err := updateExericses(&params, cc)
	if err != nil {
		return err
	}

	if len(newExercises) == 0 {
		newExercises = []*model.Exercise{}
	}

	return cc.JSON(http.StatusOK, schema.WeekDayExercise{
		Exercises: newExercises,
	})
}

// TODO: Make service
func updateExericses(param *exerciseDuplicatePostParams, cc *api.DbContext) (newExercises []*model.Exercise, err error) {
	// err = cc.ExerciseCrud.DeleteByWeekDayId(param.TimeslotId)
	// if err != nil {
	// 	return
	// }
	//
	// copyExercises, err := cc.ExerciseCrud.GetExerciseWorkSets(param.CopyTimeslotId)
	// if err != nil {
	// 	return
	// }
	//
	// newExercisesMap := make(map[int]*model.Exercise)
	// var newWorkSets []model.WorkSet
	// for _, exercise := range copyExercises {
	// 	// Adjust new exercise
	// 	exercise.ToNew(param.TimeslotId)
	//
	// 	// NOTE: Possible performance improvment, insert many
	// 	err = cc.ExerciseCrud.Insert(exercise)
	// 	if err != nil {
	// 		return
	// 	}
	// 	for _, ws := range exercise.WorkSets {
	// 		ws.ToNew(exercise.Id)
	// 		newWorkSets = append(newWorkSets, ws)
	// 	}
	// 	// Clean up
	// 	exercise.WorkSets = nil
	// 	newExercisesMap[exercise.Id] = exercise
	// }
	//
	// err = cc.WorkSetCrud.InsertMany(&newWorkSets)
	// if err != nil {
	// 	return
	// }
	//
	// for _, ws := range newWorkSets {
	// 	if e, ok := newExercisesMap[ws.ExerciseId]; ok {
	// 		e.WorkSets = append(e.WorkSets, ws)
	// 	}
	// }
	// newExercises = slices.Collect(maps.Values(newExercisesMap))
	return
}

func updateTimeslot(params *exerciseDuplicatePostParams, cc *api.DbContext) (newApiTimeslot schema.Timeslot, err error) {
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
