package crud

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/require"
)

func TestGetByUserId(t *testing.T) {
	db := testSetupDb(t)
	crud := NewExerciseType(db)
	var exerciseTypes []model.ExerciseType
	userIds := []string{"1", "2"}
	for i := range 2 {
		exerciseType := testutil.ExerciseTypeFactory(t, testutil.ExerciseTypeUserId(t, userIds[i]))
		crud.Insert(exerciseType)
		exerciseTypes = append(exerciseTypes, *exerciseType)
	}

	// Act
	dbModels, err := crud.GetByUserId(userIds[0])
	if err != nil {
		t.Fatalf("Failed to get exercise types: %v", err)
	}

	// Assert
	require.Equal(t, 1, len(dbModels))
	exerciseTypes[0].SetZeroTimes()
	dbModels[0].SetZeroTimes()
	require.Equal(t, exerciseTypes[0], dbModels[0])
}
