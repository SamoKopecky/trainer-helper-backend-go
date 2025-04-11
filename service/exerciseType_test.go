package service

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/store"

	"github.com/stretchr/testify/require"
)

func TestDuplicateDefault(t *testing.T) {
	m := store.NewMockExerciseTypeStore(t)
	service := ExerciseType{Store: m}
	mockModels := []model.ExerciseType{}
	m.EXPECT().GetByUserId("00000000-0000-0000-0000-000000000000").Return(mockModels, nil).Once()

	models, err := service.DuplicateDefault("123")

	require.Equal(t, []model.ExerciseType{}, models)
	require.Equal(t, err, ErrNoInitialExercises)

}
