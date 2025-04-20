package store

import "trainer-helper/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetExerciseWorkSets(Id int) ([]*model.Exercise, error)
	DeleteByExerciseAndTimeslot(timeslotId, exerciseId int) error
	DeleteByTimeslot(timeslotId int) error
}
