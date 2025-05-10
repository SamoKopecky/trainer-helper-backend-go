package service

import (
	"testing"
	"time"
	"trainer-helper/model"
	store "trainer-helper/store/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWeek(t *testing.T) {
	w := store.NewMockWeek(t)
	wd := store.NewMockWeekDay(t)
	service := Week{WeekStore: w, WeekDayStore: wd}
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})
	wd.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WeekDay) error {
		return nil
	})

	now := time.Now()
	newWeek := model.BuildWeek(1, now, 1, "1")
	if err := service.CreateWeek(newWeek); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	wd.Mock.AssertNumberOfCalls(t, "InsertMany", 1)
	assert.Equal(t, now.Add(time.Hour*24*0), newWeek.WeekDays[0].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*1), newWeek.WeekDays[1].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*2), newWeek.WeekDays[2].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*3), newWeek.WeekDays[3].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*4), newWeek.WeekDays[4].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*5), newWeek.WeekDays[5].DayDate)
	assert.Equal(t, now.Add(time.Hour*24*6), newWeek.WeekDays[6].DayDate)

}
