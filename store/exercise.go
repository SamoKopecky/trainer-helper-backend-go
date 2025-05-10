package store

import "trainer-helper/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetExerciseWorkSets(Id int) ([]*model.Exercise, error)
	DeleteByTimeslot(timeslotId int) error
}
