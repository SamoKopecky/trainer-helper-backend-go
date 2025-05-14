package crud

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"

	"github.com/stretchr/testify/assert"
)

func exerciseWeekDayId(weekDayId int) utils.FactoryOption[model.Exercise] {
	return func(e *model.Exercise) {
		e.WeekDayId = weekDayId
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
	weekDayId := 1
	exercise := exerciseFactory(exerciseWeekDayId(weekDayId))
	crud.Insert(exercise)
	var workSets []model.WorkSet
	for range 2 {
		ws := workSetFactory(workSetExerciseId(exercise.Id))
		workSets = append(workSets, *ws)
	}
	wsCrud.InsertMany(&workSets)

	// Act
	dbExercises, err := crud.GetExerciseWorkSets(weekDayId)
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
	var insertedExercises []*model.Exercise
	for range 3 {
		exercise := exerciseFactory(exerciseWeekDayId(1))
		crud.Insert(exercise)
		insertedExercises = append(insertedExercises, exercise)
	}
	exerciseDeleteId := insertedExercises[0].Id

	// Act
	if err := crud.Delete(exerciseDeleteId); err != nil {
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
		if exercise.Id == exerciseDeleteId {
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

func TestDeleteByWeekDayId(t *testing.T) {
	// Arrange
	db := testSetupDb(t)
	crud := NewExercise(db)
	weekDayIds := []int{1, 1, 2}
	var insertedExercises []*model.Exercise
	for _, id := range weekDayIds {
		exercise := exerciseFactory(exerciseWeekDayId(id))
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
