package service

import (
	"testing"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockGetBlocks(t *testing.T) {
	m := store.NewMockBlock(t)
	service := Block{Store: m}

	userId := "1"
	blockWithWeeks := testutil.BlockFactory(testutil.BlockUserId(userId))
	blockWithNoWeeks := testutil.BlockFactory(testutil.BlockUserId(userId))

	for j := range 2 {
		week := testutil.WeekFactory(testutil.WeekIds(userId, 0), testutil.WeekLabel(j+10))
		if j == 0 {
			for range 2 {
				weekDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, 0))
				week.WeekDays = append(week.WeekDays, *weekDay)
			}
		}
		blockWithWeeks.Weeks = append(blockWithWeeks.Weeks, *week)
	}

	mockModels := []model.Block{*blockWithWeeks, *blockWithNoWeeks}
	m.EXPECT().GetBlockWeeksByUserId("1").Return(mockModels, nil).Once()

	// Act
	actual, err := service.GetBlocks("1")

	// Assert
	require.Nil(t, err)
	assert.Equal(t, []model.WeekDay{}, actual[0].Weeks[1].WeekDays)
	assert.Equal(t, []model.Week{}, actual[1].Weeks)

}
