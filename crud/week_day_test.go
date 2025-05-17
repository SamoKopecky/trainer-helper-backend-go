package crud

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetByWeekDayId(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeekDay(db)
	var weekDays []model.WeekDay

	weekDays = append(weekDays, *testutil.WeekDayFactory(testutil.WeekDayIds("1", 1)))
	weekDays = append(weekDays, *testutil.WeekDayFactory(testutil.WeekDayIds("2", 2)))
	weekDays = append(weekDays, *testutil.WeekDayFactory(testutil.WeekDayIds("3", 2)))
	if err := crud.InsertMany(&weekDays); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	// TODO: Test with deleted
	weekDays, err := crud.GetByWeekIdWithDeleted(2)
	if err != nil {
		t.Fatalf("Failed to get week days: %v", err)
	}

	// Assert
	assert.Len(t, weekDays, 2)
	assert.Equal(t, weekDays[0].UserId, "2")
	assert.Equal(t, weekDays[1].UserId, "3")
}
