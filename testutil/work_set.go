package testutil

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WorkSetExerciseId(t *testing.T, exerciseId int) utils.FactoryOption[model.WorkSet] {
	t.Helper()
	return func(ws *model.WorkSet) {
		ws.ExerciseId = exerciseId
	}
}

func WorkSetFactory(t *testing.T, options ...utils.FactoryOption[model.WorkSet]) *model.WorkSet {
	t.Helper()
	ws := model.BuildWorkSet(utils.RandomInt(), utils.RandomInt(), nil, "10Kg")
	for _, option := range options {
		option(ws)
	}
	ws.Id = utils.RandomInt()
	return ws
}
