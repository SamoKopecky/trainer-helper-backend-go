package store

import "trainer-helper/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetExerciseWorkSets(weekDayIds []int) ([]model.Exercise, error)
	DeleteByWeekDayId(weekDayId int) error
}
