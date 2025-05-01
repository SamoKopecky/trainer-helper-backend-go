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

func WeekBlockLabel(blockLabel int) utils.FactoryOption[model.Week] {
	return func(w *model.Week) {
		w.BlockLabel = blockLabel
	}
}

func WeekFactory(options ...utils.FactoryOption[model.Week]) *model.Week {
	week := model.BuildWeek("1", time.Time{}, 0, 0, nil, nil, nil, nil, nil, nil, nil)
	for _, option := range options {
		option(week)
	}
	return week
}
