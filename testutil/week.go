package testutil

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WeekUserId(userId string) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.UserId = userId
	}
}

func WeekFactory(options ...utils.FactoryOption[model.Week]) *model.Week {
	week := model.BuildWeek(1, time.Time{}, 0, "1")
	for _, option := range options {
		option(week)
	}
	return week
}
