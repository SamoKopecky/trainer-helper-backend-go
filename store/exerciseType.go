package store

import "trainer-helper/model"

type ExerciseType interface {
	StoreBase[model.ExerciseType]
	GetByUserId(userId string) ([]model.ExerciseType, error)
}
