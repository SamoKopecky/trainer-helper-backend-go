package service

import (
	"errors"
	"trainer-helper/model"
	"trainer-helper/store"

	"github.com/google/uuid"
)

var ErrNoInitialExercises = errors.New("store: no initial exercises")

type ExerciseType struct {
	Store store.ExerciseType
}

func (et ExerciseType) DuplicateDefault(userId string) (newExerciseTypes []model.ExerciseType, err error) {
	newExerciseTypes, err = et.Store.GetByUserId(uuid.Nil.String())
	// TODO: change user ids
	if err != nil {
		return
	}
	if len(newExerciseTypes) == 0 {
		err = ErrNoInitialExercises
		return
	}

	err = et.Store.InsertMany(&newExerciseTypes)
	if err != nil {
		return
	}
	return
}
