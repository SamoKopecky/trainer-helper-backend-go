package service

import (
	"encoding/json"
	"maps"
	"slices"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/schema"
	"trainer-helper/store"
)

type AI struct {
	Fetcher           fetcher.AI
	ExerciseTypeStore store.ExerciseType
	ExerciseStore     store.Exercise
	WorkSetStore      store.WorkSet
}

func (ai AI) GenerateWeekDay(trainerId string, rawString string, weekDayId int) (newExercises []model.Exercise, err error) {
	exerciseTypes, err := ai.ExerciseTypeStore.GetByUserId(trainerId)
	if err != nil {
		return
	}
	exericseNames := make(map[string]int, len(exerciseTypes))

	for _, exerciseType := range exerciseTypes {
		exericseNames[exerciseType.Name] = exerciseType.Id
	}

	resultJson, err := ai.Fetcher.RawStringToJson(slices.Collect(maps.Keys(exericseNames)), rawString)
	if err != nil {
		return
	}

	var exercises []schema.RawExercise
	err = json.Unmarshal([]byte(resultJson), &exercises)
	if err != nil {
		return
	}

	err = ai.ExerciseStore.DeleteByWeekDayId(weekDayId)
	if err != nil {
		// TODO: raise error and stop process
		return
	}

	for i, rawExercise := range exercises {
		newExercise := model.Exercise{
			WeekDayId: weekDayId,
			Note:      &rawExercise.Note,
			GroupId:   i + 1,
		}
		if exerciseTypeId, ok := exericseNames[rawExercise.ExerciseName]; ok {
			newExercise.ExerciseTypeId = &exerciseTypeId
		}

		err = ai.ExerciseStore.Insert(&newExercise)
		if err != nil {
			return
		}
		newWorkSets := make([]model.WorkSet, len(rawExercise.WorkSets))
		for j, work_set := range rawExercise.WorkSets {
			newWorkSet := model.WorkSet{
				Intensity:  work_set.Intensity,
				ExerciseId: newExercise.Id,
				Reps:       work_set.Reps,
				Rpe:        work_set.Rpe,
			}
			newWorkSets[j] = newWorkSet
		}
		err = ai.WorkSetStore.InsertMany(&newWorkSets)
		if err != nil {
			return
		}
		newExercise.WorkSets = newWorkSets
		newExercises = append(newExercises, newExercise)
	}

	return
}
