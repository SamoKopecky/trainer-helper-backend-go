package crud

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"

	"github.com/stretchr/testify/require"
)

func timeslotTime(start, end time.Time) utils.FactoryOption[model.Timeslot] {
	return func(t *model.Timeslot) {
		t.Start = start
		t.End = end
	}
}

func timeslotIds(trainerId, traineeId string) utils.FactoryOption[model.Timeslot] {
	return func(t *model.Timeslot) {
		t.TraineeId = &traineeId
		t.TrainerId = trainerId
	}
}

func timeslotFactory(options ...utils.FactoryOption[model.Timeslot]) *model.Timeslot {
	timeslot := model.BuildTimeslot("name", time.Time{}, time.Time{}, utils.RandomUUID(), nil)
	for _, option := range options {
		option(timeslot)
	}
	return timeslot
}

func TestGetByTimeRangeAndUserId(t *testing.T) {
	// Parametrize
	var test = []struct {
		name                  string
		isTrainer             bool
		assertTimeslotIndexes []int
	}{
		{"isTrainer", true, []int{0, 1}},
		{"isNotTrainer", false, []int{0}},
	}

	// Arange
	db := testSetupDb(t)
	crud := NewTimeslot(db)
	var timeslots []*model.Timeslot
	start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	trainerIds := []string{"1", "1", "2"}
	traineeIds := []string{"1", "2", "3"}
	for i := range 3 {
		timeslot := timeslotFactory(
			timeslotTime(
				start.Add(
					time.Hour*time.Duration(i)),
				start.Add(
					time.Hour*(time.Duration(1+i)))),
			timeslotIds(trainerIds[i], traineeIds[i]))

		crud.Insert(timeslot)
		timeslots = append(timeslots, timeslot)
	}

	for _, tt := range test {
		// Act
		dbModels, err := crud.GetByTimeRangeAndUserId(start, start.Add(time.Hour*2), "1", tt.isTrainer)
		if err != nil {
			t.Fatalf("Failed to retrieve timeslots: %v", err)
		}

		// Assert
		require.Equal(t, len(tt.assertTimeslotIndexes), len(dbModels))
		var assertTimeslots []model.Timeslot
		for _, i := range tt.assertTimeslotIndexes {
			aTimeslot := *timeslots[i]
			aTimeslot.SetZeroTimes()
			assertTimeslots = append(assertTimeslots, aTimeslot)
		}
		for i := range len(dbModels) {
			dbModels[i].SetZeroTimes()
		}
		require.Equal(t, assertTimeslots, dbModels)
	}
}

func TestGetById(t *testing.T) {
	db := testSetupDb(t)
	crud := NewTimeslot(db)
	var timeslots []model.Timeslot
	for range 2 {
		timeslot := timeslotFactory()
		crud.Insert(timeslot)
		timeslots = append(timeslots, *timeslot)
	}

	// Act
	dbModel, err := crud.GetById(timeslots[0].Id)
	if err != nil {
		t.Fatalf("Failed to retrieve timeslot: %v", err)
	}

	// Assert
	aTimeslot := timeslots[0]
	aTimeslot.SetZeroTimes()
	dbModel.SetZeroTimes()

	require.Equal(t, aTimeslot, dbModel)
}
