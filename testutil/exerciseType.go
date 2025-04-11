package testutil

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func ExerciseTypeUserId(t *testing.T, userId string) utils.FactoryOption[model.ExerciseType] {
	t.Helper()
	return func(et *model.ExerciseType) {
		et.UserId = userId
	}
}
func ExerciseTypeFactory(t *testing.T, options ...utils.FactoryOption[model.ExerciseType]) *model.ExerciseType {
	t.Helper()
	exerciseType := model.BuildExerciseType(utils.RandomUUID(), "name", nil, nil, nil)
	for _, option := range options {
		option(exerciseType)
	}
	return exerciseType

}
