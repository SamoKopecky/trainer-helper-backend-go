package service

import (
	"encoding/json"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/schema"
	"trainer-helper/store"
	"trainer-helper/utils"
)

type AI struct {
	Fetcher           fetcher.AI
	ExerciseTypeStore store.ExerciseType
	ExerciseStore     store.Exercise
	WorkSetStore      store.WorkSet
}

func (ai AI) GenerateWeekDay(trainerId string, rawString string, weekDayId int) error {
	exerciseTypes, err := ai.ExerciseTypeStore.GetByUserId(trainerId)
	if err != nil {
		return err
	}
	exericseNames := make([]string, len(exerciseTypes))

	for i := range exerciseTypes {
		exericseNames[i] = exerciseTypes[i].Name
	}

	resultJson, err := ai.Fetcher.RawStringToJson(exericseNames, rawString)
	if err != nil {
		return err
	}

	var exercises []schema.RawExercise
	err = json.Unmarshal([]byte(resultJson), &exercises)
	if err != nil {
		return err
	}

	utils.PrettyPrint(exercises)

	err = ai.ExerciseStore.DeleteByWeekDayId(weekDayId)
	if err != nil {
		// TODO: raise error and stop process
		return err
	}

	for i, exercise := range exercises {
		newExercise := model.Exercise{
			WeekDayId: weekDayId,
			Note:      &exercise.Note,
			GroupId:   i + 1,
		}
		err := ai.ExerciseStore.Insert(&newExercise)
		if err != nil {
			return err
		}
		newWorkSets := make([]model.WorkSet, len(exercise.WorkSets))
		for _, work_set := range exercise.WorkSets {
			newWorkSet := model.WorkSet{
				Intensity:  work_set.Intensity,
				ExerciseId: newExercise.Id,
				Reps:       work_set.Reps,
				Rpe:        work_set.Rpe,
			}
			newWorkSets = append(newWorkSets, newWorkSet)
		}
		err = ai.WorkSetStore.InsertMany(&newWorkSets)
		if err != nil {
			return err
		}
	}

	return nil
}
