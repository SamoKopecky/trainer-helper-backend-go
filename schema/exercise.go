package schema

import "trainer-helper/model"

type WeekDayExercise struct {
	Exercises []*model.Exercise `json:"exercises"`
}
