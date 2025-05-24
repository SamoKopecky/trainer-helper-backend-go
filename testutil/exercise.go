package testutil

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func ExerciseWeekDayId(t *testing.T, weekDayId int) utils.FactoryOption[model.Exercise] {
	t.Helper()
	return func(e *model.Exercise) {
		e.WeekDayId = weekDayId
	}
}

func ExerciseFactory(t *testing.T, options ...utils.FactoryOption[model.Exercise]) *model.Exercise {
	t.Helper()
	note := "note"
	exercise := model.BuildExercise(utils.RandomInt(), utils.RandomInt(), &note, nil)
	for _, option := range options {
		option(exercise)
	}
	exercise.Id = utils.RandomInt()

	return exercise
}
