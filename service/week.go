package service

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/store"
	"trainer-helper/utils"
)

const DaysInAWeek = 7

type Week struct {
	WeekStore    store.Week
	WeekDayStore store.WeekDay
}

func (w Week) CreateWeek(newWeek *model.Week) (err error) {
	lastDate, err := w.WeekStore.GetLastWeekDate(newWeek.BlockId)
	if err != nil {
		return
	}

	if lastDate.IsZero() {
		// If no week exists, just get next monday
		newWeek.StartDate = utils.GetNextMonday(time.Now())
	} else {
		// If some week already exists, add 7 days to get next monday
		newWeek.StartDate = utils.GetNextMonday(lastDate).AddDate(0, 0, 7)
	}

	if err = w.WeekStore.Insert(newWeek); err != nil {
		return
	}

	newWeek.WeekDays = make([]model.WeekDay, DaysInAWeek)
	for i := range DaysInAWeek {
		newDate := newWeek.StartDate.AddDate(0, 0, i)
		newWeek.WeekDays[i] = *model.BuildWeekDay(newWeek.Id, newWeek.UserId, newDate, nil)
	}

	if err = w.WeekDayStore.InsertMany(&newWeek.WeekDays); err != nil {
		return
	}

	return nil

}
