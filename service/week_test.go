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

func TestCreateWeek_NotFirst(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := time.Date(2025, 05, 12, 0, 0, 0, 0, time.UTC)
	service := Week{WeekStore: w}
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(nextMonday, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, nextMonday, 1, "1")
	// Act
	if err := service.CreateWeek(newWeek, false); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// +7 to get to the next monday
	assert.Equal(t, nextMonday.AddDate(0, 0, 7), newWeek.StartDate)
}

func TestCreateWeek__NoLastWeek__NotFirst(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := utils.GetNextMonday(time.Now())
	service := Week{WeekStore: w}
	// Rerturn 0 time
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(time.Time{}, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, time.Time{}, 1, "1")
	// Act
	if err := service.CreateWeek(newWeek, false); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// 7 to get to the next monday
	assert.Equal(t, nextMonday.Day(), newWeek.StartDate.Day())
}

func TestCreateWeek__First(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := time.Date(2025, 05, 12, 0, 0, 0, 0, time.UTC)
	service := Week{WeekStore: w}
	w.EXPECT().GetPreviousBlockId(mock.Anything).Return(nextMonday, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, nextMonday, 1, "1")
	// Act
	if err := service.CreateWeek(newWeek, true); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// +7 to get to the next monday
	assert.Equal(t, nextMonday.AddDate(0, 0, 7), newWeek.StartDate)
}
