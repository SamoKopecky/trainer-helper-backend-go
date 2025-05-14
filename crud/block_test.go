package crud

import (
	"strconv"
	"testing"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func zeroTimestamps(block *model.Block) {
	if block == nil {
		return
	}
	for i := range block.Weeks {
		block.Weeks[i].SetZeroTimes()
	}
}

func TestGetBlockWeeksByUserId(t *testing.T) {
	db := testSetupDb(t)
	blockCrud := NewBlock(db)
	weekCrud := NewWeek(db)
	weekDayCrud := NewWeekDay(db)

	var blocks []model.Block
	var weeks []model.Week
	var weekDays []model.WeekDay

	for i := range 2 {
		userId := strconv.Itoa(i)
		block := testutil.BlockFactory(testutil.BlockUserId(userId))
		blockCrud.Insert(block)
		blocks = append(blocks, *block)

		for j := range 2 {
			week := testutil.WeekFactory(testutil.WeekIds(userId, block.Id))
			weekCrud.Insert(week)
			weeks = append(weeks, *week)

			if j == 0 {
				for range 2 {
					weekDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, week.Id))
					weekDayCrud.Insert(weekDay)
					weekDays = append(weekDays, *weekDay)
				}
			}
		}
	}

	// Act
	fetchedBlocks, err := blockCrud.GetBlockWeeksByUserId(blocks[0].UserId)
	if err != nil {
		t.Fatalf("Failed to retrieve weeks: %v", err)
	}
	fetchedBlock := fetchedBlocks[0]
	zeroTimestamps(&fetchedBlock)

	// Assert
	require.Len(t, fetchedBlocks, 1)
	require.Len(t, fetchedBlocks[0].Weeks, 2)

	exepectedBlock := blocks[0]
	exepectedBlock.Weeks = append(exepectedBlock.Weeks, weeks[0])
	exepectedBlock.Weeks = append(exepectedBlock.Weeks, weeks[1])
	zeroTimestamps(&exepectedBlock)

	assert.Equal(t, exepectedBlock, fetchedBlock)

}
