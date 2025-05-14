package testutil

import (
	"trainer-helper/model"
	"trainer-helper/utils"
)

func ExerciseWeekDayId(weekDayId int) utils.FactoryOption[model.Exercise] {
	return func(e *model.Exercise) {
		e.WeekDayId = weekDayId
	}
}

func ExerciseFactory(options ...utils.FactoryOption[model.Exercise]) *model.Exercise {
	note := "note"
	exercise := model.BuildExercise(utils.RandomInt(), utils.RandomInt(), &note, nil)
	for _, option := range options {
		option(exercise)
	}
	return exercise
}
