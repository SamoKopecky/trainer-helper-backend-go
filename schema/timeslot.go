package schema

import "trainer-helper/model"

type Timeslot struct {
	model.Timeslot
	User    *model.User    `json:"user"`
	WeekDay *model.WeekDay `json:"week_day"`
}
