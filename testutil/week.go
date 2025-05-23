package testutil

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WeekIds(userId string, blockId int) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.UserId = userId
		w.BlockId = blockId
	}
}

func WeekLabel(label int) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.Label = label
	}
}

func WeekDate(date time.Time) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.StartDate = date
	}
}

func WeekId(id int) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.Id = id
	}
}

func WeekFactory(options ...utils.FactoryOption[model.Week]) *model.Week {
	week := model.BuildWeek(1, time.Now(), 0, "1")
	week.Id = utils.RandomInt()
	for _, option := range options {
		option(week)
	}
	return week
}
