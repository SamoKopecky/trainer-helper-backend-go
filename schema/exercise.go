package schema

import "trainer-helper/model"

type WeekDayExercise struct {
	WeekDay   WeekDay           `json:"week_day"`
	Exercises []*model.Exercise `json:"exercises"`
}
