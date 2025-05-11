package crud

import (
	"testing"
	"time"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGetLastWeekDate(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeek(db)
	var weeks []model.Week

	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 1),
			testutil.WeekLabel(2),
			testutil.WeekDate(time.Now().Add(time.Hour*2))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(3),
			testutil.WeekDate(time.Now().Add(time.Hour*3))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(1),
			testutil.WeekDate(time.Now().Add(time.Hour*1))))
	if err := crud.InsertMany(&weeks); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	startDate, err := crud.GetLastWeekDate(2)
	if err != nil {
		t.Fatalf("Failed to get week date: %v", err)
	}

	// Assert
	assert.Equal(t, time.Now().Add(time.Hour*3).UTC().Hour(), startDate.UTC().Hour())
}
