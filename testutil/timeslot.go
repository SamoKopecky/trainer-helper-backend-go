package testutil

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func TimeslotTime(t *testing.T, start, end time.Time) utils.FactoryOption[model.Timeslot] {
	t.Helper()
	return func(t *model.Timeslot) {
		t.Start = start
		t.End = end
	}
}

func TimeslotIds(t *testing.T, trainerId, traineeId string) utils.FactoryOption[model.Timeslot] {
	t.Helper()
	return func(t *model.Timeslot) {
		t.TraineeId = &traineeId
		t.TrainerId = trainerId
	}
}

func TimeslotFactory(t *testing.T, options ...utils.FactoryOption[model.Timeslot]) *model.Timeslot {
	t.Helper()
	timeslot := model.BuildTimeslot(time.Time{}, time.Time{}, utils.RandomUUID(), nil)
	for _, option := range options {
		option(timeslot)
	}
	timeslot.Id = utils.RandomInt()
	return timeslot
}
