package service

import (
	"sort"
	"trainer-helper/model"
	"trainer-helper/store"
)

type Exercise struct {
	Store store.Exercise
}

func (e Exercise) GetExerciseWorkSets(weekDayIds []int) (exercisesMap map[int][]model.Exercise, err error) {
	exercisesMap = make(map[int][]model.Exercise)
	rawExercises, err := e.Store.GetExerciseWorkSets(weekDayIds)
	if err != nil || rawExercises == nil {
		return
	}

	sort.Slice(rawExercises, func(i, j int) bool {
		if rawExercises[i].GroupId == rawExercises[j].GroupId {
			return rawExercises[i].Id < rawExercises[j].Id
		}
		return rawExercises[i].GroupId < rawExercises[j].GroupId
	})
	for _, exercise := range rawExercises {
		exercise.SortWorkSets()
	}

	for i, exercise := range rawExercises {
		if len(rawExercises[i].WorkSets) == 0 {
			rawExercises[i].WorkSets = []model.WorkSet{}
		}

		exercisesMap[exercise.WeekDayId] = append(exercisesMap[exercise.WeekDayId], exercise)
	}

	return
}
