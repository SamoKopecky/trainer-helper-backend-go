package crud

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"

	"github.com/stretchr/testify/assert"
)

func exerciseTimeslotId(timeslotId int) utils.FactoryOption[model.Exercise] {
	return func(e *model.Exercise) {
		e.TimeslotId = timeslotId
	}
}

func exerciseFactory(options ...utils.FactoryOption[model.Exercise]) *model.Exercise {
	note := "note"
	exercise := model.BuildExercise(utils.RandomInt(), utils.RandomInt(), &note, nil)
	for _, option := range options {
		option(exercise)
	}
	return exercise
}

func TestGetExerciseWorkSets(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	wsCrud := NewWorkSet(db)
	timeslotId := 1
	exercise := exerciseFactory(exerciseTimeslotId(timeslotId))
	crud.Insert(exercise)
	var workSets []model.WorkSet
	for range 2 {
		ws := workSetFactory(workSetExerciseId(exercise.Id))
		workSets = append(workSets, *ws)
	}
	wsCrud.InsertMany(&workSets)

	// Act
	dbExercises, err := crud.GetExerciseWorkSets(timeslotId)
	if err != nil {
		t.Fatalf("Failed to get exercises: %v", err)
	}

	// Assert
	assert.Equal(t, 1, len(dbExercises))
	exercise.WorkSets = workSets
	exercise.SetZeroTimes()
	for i := range len(workSets) {
		workSets[i].SetZeroTimes()
	}
	dbExercise := dbExercises[0]
	dbExercise.SetZeroTimes()
	for i := range len(dbExercise.WorkSets) {
		dbExercise.WorkSets[i].SetZeroTimes()
	}

	assert.Equal(t, dbExercise, exercise)
}

func TestDeleteByExerciseAndTimeslot(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	timeslotIds := []int{1, 1, 2}
	var insertedExercises []*model.Exercise
	for _, id := range timeslotIds {
		exercise := exerciseFactory(exerciseTimeslotId(id))
		crud.Insert(exercise)
		insertedExercises = append(insertedExercises, exercise)
	}
	exerciseDeleteId := insertedExercises[0].Id
	timeslotDeleteId := timeslotIds[0]

	// Act
	if err := crud.DeleteByExerciseAndTimeslot(timeslotDeleteId, exerciseDeleteId); err != nil {
		t.Fatalf("Failed to delete exercises: %v", err)
	}

	// Assert
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Equal(t, len(insertedExercises)-1, len(dbModels))
	var assertExercises []model.Exercise
	for i := range len(insertedExercises) {
		exercise := insertedExercises[i]
		if exercise.Id == exerciseDeleteId && exercise.TimeslotId == timeslotDeleteId {
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

func TestDeleteByTimeslot(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	timeslotIds := []int{1, 1, 2}
	var insertedExercises []*model.Exercise
	for _, id := range timeslotIds {
		exercise := exerciseFactory(exerciseTimeslotId(id))
		crud.Insert(exercise)
		insertedExercises = append(insertedExercises, exercise)
	}
	timeslotDeleteId := timeslotIds[0]

	// Act
	if err := crud.DeleteByTimeslot(timeslotDeleteId); err != nil {
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
		if exercise.TimeslotId == timeslotDeleteId {
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
