package testutil

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WeekDayIds(t *testing.T, userId string, weekId int) utils.FactoryOption[model.WeekDay] {
	t.Helper()
	return func(wd *model.WeekDay) {
		wd.UserId = userId
		wd.WeekId = weekId
	}
}

func WeekDayTime(t *testing.T, date time.Time) utils.FactoryOption[model.WeekDay] {
	t.Helper()
	return func(wd *model.WeekDay) {
		wd.DayDate = date
	}
}

func WeekDayFactory(t *testing.T, options ...utils.FactoryOption[model.WeekDay]) *model.WeekDay {
	t.Helper()
	now := time.Now()
	weekDay := &model.WeekDay{
		UserId:  "1",
		WeekId:  utils.RandomInt(),
		DayDate: now,
		Name:    nil,
	}
	weekDay.Id = utils.RandomInt()

	for _, option := range options {
		option(weekDay)
	}
	return weekDay
}
