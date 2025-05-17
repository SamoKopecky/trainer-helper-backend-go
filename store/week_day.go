package store

import "trainer-helper/model"

type WeekDay interface {
	StoreBase[model.WeekDay]
	GetByWeekIdWithDeleted(weekId int) (weekDays []model.WeekDay, err error)
	GetByTimeslotIds(timeslotIds []int) (weekDays []model.WeekDay, err error)
}
