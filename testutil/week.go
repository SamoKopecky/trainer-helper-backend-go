package testutil

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func WeekIds(t *testing.T, userId string, blockId int) utils.FactoryOption[model.Week] {
	t.Helper()
	return func(w *model.Week) {
		w.UserId = userId
		w.BlockId = blockId
	}
}

func WeekLabel(t *testing.T, label int) utils.FactoryOption[model.Week] {
	t.Helper()
	return func(w *model.Week) {
		w.Label = label
	}
}

func WeekDate(t *testing.T, date time.Time) utils.FactoryOption[model.Week] {
	t.Helper()
	return func(w *model.Week) {
		w.StartDate = date
	}
}

func WeekId(t *testing.T, id int) utils.FactoryOption[model.Week] {
	t.Helper()
	return func(w *model.Week) {
		w.Id = id
	}
}

func WeekFactory(t *testing.T, options ...utils.FactoryOption[model.Week]) *model.Week {
	t.Helper()
	week := model.BuildWeek(1, time.Now(), 0, "1", nil)
	week.Id = utils.RandomInt()
	for _, option := range options {
		option(week)
	}
	return week
}
