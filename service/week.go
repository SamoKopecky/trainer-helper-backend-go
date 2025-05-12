package service

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/store"
	"trainer-helper/utils"
)

const DaysInAWeek = 7

type Week struct {
	WeekStore store.Week
}

func (w Week) CreateWeek(newWeek *model.Week, isFirst bool) (err error) {
	var lastDate time.Time
	if isFirst {
		lastDate, err = w.WeekStore.GetPreviousBlockId(newWeek.UserId)
	} else {
		lastDate, err = w.WeekStore.GetLastWeekDate(newWeek.BlockId)
	}
	if err != nil {
		return err
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
	newWeek.WeekDays = []model.WeekDay{}

	return nil

}
