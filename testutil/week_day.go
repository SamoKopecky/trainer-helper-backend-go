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

func WeekDayFactory(options ...utils.FactoryOption[model.WeekDay]) *model.WeekDay {
	now := time.Now()
	weekDay := &model.WeekDay{
		UserId:  "1",
		WeekId:  utils.RandomInt(),
		DayDate: now,
		Name:    "name",
	}

	for _, option := range options {
		option(weekDay)
	}
	return weekDay
}
