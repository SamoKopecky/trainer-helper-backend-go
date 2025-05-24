package crud

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetByWeekDayId(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeekDay(db)
	var weekDays []model.WeekDay

	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "1", 1)))
	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "2", 2)))
	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "3", 2)))
	now := time.Now()
	weekDays[0].DeletedAt = &now
	if err := crud.InsertMany(&weekDays); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	weekDays, err := crud.GetByWeekIdWithDeleted(2)
	if err != nil {
		t.Fatalf("Failed to get week days: %v", err)
	}

	// Assert
	assert.Len(t, weekDays, 2)
	assert.Equal(t, weekDays[0].UserId, "2")
	assert.Equal(t, weekDays[1].UserId, "3")
}

func TestGetByDate(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeekDay(db)
	var weekDays []model.WeekDay
	now := time.Now()

	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "2", 1), testutil.WeekDayTime(t, now)))
	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "1", 1), testutil.WeekDayTime(t, now.AddDate(0, 0, 1))))
	weekDays = append(weekDays, *testutil.WeekDayFactory(t, testutil.WeekDayIds(t, "1", 1), testutil.WeekDayTime(t, now)))
	if err := crud.InsertMany(&weekDays); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	dbWeekDays, err := crud.GetByDate(now, "1")
	if err != nil {
		t.Fatalf("Failed to get week day: %v", err)
	}

	// Asser
	assert.Len(t, dbWeekDays, 1)
	assert.Equal(t, dbWeekDays[0].Id, weekDays[2].Id)

}
