package testutil

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WeekDayIds(userId string, weekId int) utils.FactoryOption[model.WeekDay] {
	return func(wd *model.WeekDay) {
		wd.UserId = userId
		wd.WeekId = weekId
	}
}

func WeekDayTime(date time.Time) utils.FactoryOption[model.WeekDay] {
	return func(wd *model.WeekDay) {
		wd.DayDate = date
	}
}

func WeekDayFactory(options ...utils.FactoryOption[model.WeekDay]) *model.WeekDay {
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
