package crud

import (
	"slices"
	"testing"
	"trainer-helper/model"

	"github.com/stretchr/testify/assert"
)

func workSetExerciseId(exerciseId int) FactoryOption[model.WorkSet] {
	return func(ws *model.WorkSet) {
		ws.ExerciseId = exerciseId
	}
}
func workSetFactory(options ...FactoryOption[model.WorkSet]) *model.WorkSet {
	rpe := randomInt()
	ws := model.BuildWorkSet(randomInt(), randomInt(), &rpe, "10Kg")
	for _, option := range options {
		option(ws)
	}
	return ws
}

func TestInsertMany(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)

	// Arange
	var workSets []model.WorkSet
	for range 2 {
		workSets = append(workSets, *workSetFactory())
	}

	// Act
	if err := crud.InsertMany(&workSets); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Assert
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}

	for i := range workSets {
		workSets[i].Timestamp.SetZeroTimes()
		dbModels[i].Timestamp.SetZeroTimes()
	}

	assert.EqualValues(t, dbModels, workSets, "Work sets should be equal")
}

func TestDeleteMany(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWorkSet(db)

	// Arange
	var workSets []model.WorkSet
	for range 3 {
		workSets = append(workSets, *workSetFactory())
	}
	if err := crud.InsertMany(&workSets); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}
	toDelete := []int{workSets[0].Id}
	assert.Equal(t, 3, len(workSets))

	// Act
	deleted, err := crud.DeleteMany(toDelete)
	if err != nil {
		t.Fatalf("Failed to delete work sets: %v", err)
	}

	// Asert
	assert.Equal(t, 1, deleted)
	dbModels, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Equal(t, 2, len(dbModels))

	var toAssert []model.WorkSet
	for i := range workSets {
		if !slices.Contains(toDelete, workSets[i].Id) {
			workSets[i].Timestamp.SetZeroTimes()
			toAssert = append(toAssert, workSets[i])
		}
	}
	for i := range dbModels {
		dbModels[i].Timestamp.SetZeroTimes()
	}

	assert.EqualValues(t, toAssert, dbModels, "Work sets should be equal")
}
