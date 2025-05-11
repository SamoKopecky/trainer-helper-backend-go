package service

import (
	"testing"
	"time"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWeek(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	wd := store.NewMockWeekDay(t)
	now := time.Date(2025, 05, 10, 0, 0, 0, 0, time.UTC)
	service := Week{WeekStore: w, WeekDayStore: wd}
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(now.AddDate(0, 0, 10), nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})
	wd.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WeekDay) error {
		return nil
	})

	newWeek := model.BuildWeek(1, now, 1, "1")
	// Act
	if err := service.CreateWeek(newWeek); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	wd.Mock.AssertNumberOfCalls(t, "InsertMany", 1)
	// 10 offset + 6 to get to the next monday
	assert.Equal(t, now.AddDate(0, 0, 16), newWeek.StartDate)
	assert.Equal(t, now.AddDate(0, 0, 16), newWeek.WeekDays[0].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 17), newWeek.WeekDays[1].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 18), newWeek.WeekDays[2].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 19), newWeek.WeekDays[3].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 20), newWeek.WeekDays[4].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 21), newWeek.WeekDays[5].DayDate)
	assert.Equal(t, now.AddDate(0, 0, 22), newWeek.WeekDays[6].DayDate)
}

func TestCreateWeek__NoLastWeek(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	wd := store.NewMockWeekDay(t)
	now := time.Now()
	nextMonday := utils.GetNextMonday(now)
	service := Week{WeekStore: w, WeekDayStore: wd}
	// Rerturn 0 time
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(time.Time{}, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})
	wd.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WeekDay) error {
		return nil
	})

	newWeek := model.BuildWeek(1, now, 1, "1")
	// Act
	if err := service.CreateWeek(newWeek); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	wd.Mock.AssertNumberOfCalls(t, "InsertMany", 1)
	// 10 offset + 6 to get to the next monday
	assert.Equal(t, nextMonday.Day(), newWeek.StartDate.Day())
}
