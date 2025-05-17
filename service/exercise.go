package service

import (
	"sort"
	"trainer-helper/model"
	"trainer-helper/store"
)

type Exercise struct {
	Store store.Exercise
}

// Get maped exercises by week day id
func (e Exercise) GetExerciseWorkSets(weekDayIds []int) (exercises []model.Exercise, err error) {
	exercises, err = e.Store.GetExerciseWorkSets(weekDayIds)
	if err != nil || exercises == nil {
		return []model.Exercise{}, nil
	}

	e.sortExercises(exercises)

	for i := range exercises {
		if len(exercises[i].WorkSets) == 0 {
			exercises[i].WorkSets = []model.WorkSet{}
		}

	}

	return
}

func (e Exercise) sortExercises(rawExercises []model.Exercise) {
	sort.Slice(rawExercises, func(i, j int) bool {
		if rawExercises[i].GroupId == rawExercises[j].GroupId {
			return rawExercises[i].Id < rawExercises[j].Id
		}
		return rawExercises[i].GroupId < rawExercises[j].GroupId
	})
	for _, exercise := range rawExercises {
		exercise.SortWorkSets()
	}
}
