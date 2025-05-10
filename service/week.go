package service

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/store"
)

const DaysInAWeek = 7

type Week struct {
	WeekStore    store.Week
	WeekDayStore store.WeekDay
}

func (w Week) CreateWeek(newWeek *model.Week) (err error) {
	if err = w.WeekStore.Insert(newWeek); err != nil {
		return
	}

	newWeek.WeekDays = make([]model.WeekDay, DaysInAWeek)
	for i := range DaysInAWeek {
		newDate := newWeek.StartDate.Add(time.Hour * 24 * time.Duration(i))
		newWeek.WeekDays[i] = *model.BuildWeekDay(newWeek.Id, newWeek.UserId, newDate, nil)
	}

	if err = w.WeekDayStore.InsertMany(&newWeek.WeekDays); err != nil {
		return
	}

	return nil

}
