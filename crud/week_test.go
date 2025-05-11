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
			testutil.WeekDate(time.Now().AddDate(0, 0, 2))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(3),
			testutil.WeekDate(time.Now().AddDate(0, 0, 3))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(1),
			testutil.WeekDate(time.Now().AddDate(0, 0, 1))))
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
		*testutil.BlockFactory(testutil.BlockUserId("1"),
			testutil.BlockId(1),
			testutil.BlockLabel(1),
		))
	blocks = append(blocks,
		*testutil.BlockFactory(testutil.BlockUserId("1"),
			testutil.BlockId(2),
			testutil.BlockLabel(2),
		))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 1),
			testutil.WeekLabel(1),
			testutil.WeekDate(time.Now().AddDate(0, 0, 1))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(1),
			testutil.WeekDate(time.Now().AddDate(0, 0, 2))))
	weeks = append(weeks,
		*testutil.WeekFactory(testutil.WeekIds("1", 2),
			testutil.WeekLabel(2),
			testutil.WeekDate(time.Now().AddDate(0, 0, 3))))
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
