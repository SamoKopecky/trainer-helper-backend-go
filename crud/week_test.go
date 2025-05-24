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
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 1),
			testutil.WeekLabel(t, 2),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 2))))
	weeks = append(weeks,
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 2),
			testutil.WeekLabel(t, 3),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 3))))
	weeks = append(weeks,
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 2),
			testutil.WeekLabel(t, 1),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 1))))
	if err := crud.InsertMany(&weeks); err != nil {
		t.Fatalf("Failed to insert work sets: %v", err)
	}

	// Act
	startDate, err := crud.GetLastWeekDate(2)
	if err != nil {
		t.Fatalf("Failed to get week date: %v", err)
	}

	// Assert
	assert.Equal(t, time.Now().AddDate(0, 0, 3).Day(), startDate.UTC().Day())
}

func TestGetPreviousBlockId(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeek(db)
	blockCrud := NewBlock(db)
	var weeks []model.Week
	var blocks []model.Block

	// 2 Blocks (1, 2)
	// Block 1 has 1 week (1)
	// Block 2 has 2 weeks (1 , 2)
	blocks = append(blocks,
		*testutil.BlockFactory(t, testutil.BlockUserId(t, "1"),
			testutil.BlockId(t, 1),
			testutil.BlockLabel(t, 1),
		))
	blocks = append(blocks,
		*testutil.BlockFactory(t, testutil.BlockUserId(t, "1"),
			testutil.BlockId(t, 2),
			testutil.BlockLabel(t, 2),
		))
	weeks = append(weeks,
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 1),
			testutil.WeekLabel(t, 1),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 1))))
	weeks = append(weeks,
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 2),
			testutil.WeekLabel(t, 1),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 2))))
	weeks = append(weeks,
		*testutil.WeekFactory(t, testutil.WeekIds(t, "1", 2),
			testutil.WeekLabel(t, 2),
			testutil.WeekDate(t, time.Now().AddDate(0, 0, 3))))
	if err := crud.InsertMany(&weeks); err != nil {
		t.Fatalf("Failed to insert weeks: %v", err)
	}
	if err := blockCrud.InsertMany(&blocks); err != nil {
		t.Fatalf("Failed to insert blocks: %v", err)
	}

	// Act
	startDate, err := crud.GetPreviousBlockId("1")
	if err != nil {
		t.Fatalf("Failed to get week date: %v", err)
	}

	// Assert
	assert.Equal(t, time.Now().AddDate(0, 0, 3).Day(), startDate.UTC().Day())
}

func TestGetClosestToDate(t *testing.T) {
	testCases := []struct {
		name       string
		dayOffset  int
		expectedId int
	}{
		{
			name:       "exact match",
			dayOffset:  2,
			expectedId: 1,
		},
		{
			name:       "greater",
			dayOffset:  4,
			expectedId: 2,
		},
		{
			name:       "smaller",
			dayOffset:  0,
			expectedId: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := testSetupDb(t)
			crud := NewWeek(db)
			var weeks []model.Week

			weeks = append(weeks,
				*testutil.WeekFactory(t, testutil.WeekId(t, 1),
					testutil.WeekIds(t, "1", 1),
					testutil.WeekDate(t, time.Now().AddDate(0, 0, 2))))
			weeks = append(weeks,
				*testutil.WeekFactory(t, testutil.WeekId(t, 2),
					testutil.WeekIds(t, "1", 1),
					testutil.WeekDate(t, time.Now().AddDate(0, 0, 3))))
			weeks = append(weeks,
				*testutil.WeekFactory(t, testutil.WeekId(t, 3),
					testutil.WeekIds(t, "1", 1),
					testutil.WeekDate(t, time.Now().AddDate(0, 0, 1))))
			if err := crud.InsertMany(&weeks); err != nil {
				t.Fatalf("Failed to insert work sets: %v", err)
			}

			// Act
			closestWeek, err := crud.GetClosestToDate(time.Now().AddDate(0, 0, tc.dayOffset), "1")
			if err != nil {
				t.Fatalf("Failed to get week date: %v", err)
			}

			// Assert
			assert.Equal(t, tc.expectedId, closestWeek.Id)

		})
	}
}
