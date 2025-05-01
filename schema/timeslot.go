package schema

import "trainer-helper/model"

type Timeslot struct {
	model.Timeslot
	UserName     *string `json:"user_name"`
	UserNickname *string `json:"user_nickname"`
}

type TimeslotExercises struct {
	Timeslot  Timeslot          `json:"timeslot"`
	Exercises []*model.Exercise `json:"exercises"`
}
