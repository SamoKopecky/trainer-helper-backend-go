package store

import (
	"time"
	"trainer-helper/model"
)

type WeekDay interface {
	StoreBase[model.WeekDay]
	GetByWeekIdsWithDeleted(weekIds []int) (weekDays []model.WeekDay, err error)
	GetByTimeslotIds(timeslotIds []int) (weekDays []model.WeekDay, err error)
	GetByDate(dayDate time.Time, userId string) (weekDays []model.WeekDay, err error)
	DeleteTimeslot(weekId int) error
}
