package service

import (
	"testing"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/testutil"

	"slices"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateDefault(t *testing.T) {
	m := store.NewMockExerciseType(t)
	service := ExerciseType{Store: m}
	mockModels := []model.ExerciseType{*testutil.ExerciseTypeFactory(t)}
	assertModels := slices.Clone(mockModels)
	assertModels[0].UserId = "123"
	m.EXPECT().GetByUserId("00000000-0000-0000-0000-000000000000").Return(mockModels, nil).Once()
	m.EXPECT().InsertMany(&assertModels).Return(nil).Once()

	// Act
	models, err := service.DuplicateDefault("123")

	// Assert
	assert.Equal(t, assertModels, models)
	assert.Equal(t, err, nil)

}

func TestDuplicateDefault_NoInitExerciseTypes(t *testing.T) {
	m := store.NewMockExerciseType(t)
	service := ExerciseType{Store: m}
	mockModels := []model.ExerciseType{}
	m.EXPECT().GetByUserId("00000000-0000-0000-0000-000000000000").Return(mockModels, nil).Once()

	// Act
	models, err := service.DuplicateDefault("123")

	// Assert
	assert.Equal(t, models, mockModels)
	assert.Equal(t, err, ErrNoInitialExercises)

}
