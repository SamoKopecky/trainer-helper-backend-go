package schema

import "trainer-helper/model"

type TimeslotExercises struct {
	Timeslot  Timeslot          `json:"timeslot"`
	Exercises []*model.Exercise `json:"exercises"`
}
