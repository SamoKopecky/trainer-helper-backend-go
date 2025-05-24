package crud

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
)

func setZeroTimes(exercises []model.Exercise) {
	exercises[0].SetZeroTimes()
	exercises[1].SetZeroTimes()
	for i := range len(exercises[0].WorkSets) {
		exercises[0].WorkSets[i].SetZeroTimes()
	}
}

func TestGetExerciseWorkSets(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	wsCrud := NewWorkSet(db)
	exercise1 := testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, 1))
	exercise2 := testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, 2))
	exercise3 := testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, 3))
	exercises := []model.Exercise{*exercise1, *exercise2, *exercise3}
	crud.InsertMany(&exercises)

	var workSets []model.WorkSet
	for range 2 {
		ws := testutil.WorkSetFactory(t, testutil.WorkSetExerciseId(t, exercises[0].Id))
		workSets = append(workSets, *ws)
	}
	wsCrud.InsertMany(&workSets)

	// Act
	dbExercises, err := crud.GetExerciseWorkSets([]int{1, 2})
	if err != nil {
		t.Fatalf("Failed to get exercises: %v", err)
	}

	// Assert
	assert.Len(t, dbExercises, 2)
	expectedExercises := exercises[:2]
	expectedExercises[0].WorkSets = workSets
	setZeroTimes(expectedExercises)
	setZeroTimes(dbExercises)

	assert.Equal(t, dbExercises, expectedExercises)
}

func TestDeleteByWeekDayId(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	weekDayIds := []int{1, 1, 2}
	var insertedExercises []*model.Exercise
	for _, id := range weekDayIds {
		exercise := testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, id))
		crud.Insert(exercise)
		insertedExercises = append(insertedExercises, exercise)
	}
	weekDayDeleteId := weekDayIds[0]

	// Act
	if err := crud.DeleteByWeekDayId(weekDayDeleteId); err != nil {
		t.Fatalf("Failed to delete exercises: %v", err)
	}

	// Assert
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Equal(t, len(insertedExercises)-2, len(dbModels))
	var assertExercises []model.Exercise
	for i := range len(insertedExercises) {
		exercise := insertedExercises[i]
		if exercise.WeekDayId == weekDayDeleteId {
			continue
		}
		exercise.SetZeroTimes()
		assertExercises = append(assertExercises, *exercise)
	}
	for i := range len(dbModels) {
		dbModels[i].SetZeroTimes()
	}

	assert.Equal(t, dbModels, assertExercises)
}
