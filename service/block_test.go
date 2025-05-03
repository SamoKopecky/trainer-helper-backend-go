package service

import (
	"testing"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/testutil"
	"trainer-helper/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWeekGetBlocks(t *testing.T) {
	m := store.NewMockBlock(t)
	service := Block{Store: m}

	userId := "1"
	block := testutil.BlockFactory(testutil.BlockUserId(userId))

	for j := range 2 {
		week := testutil.WeekFactory(testutil.WeekIds(userId, 0), testutil.WeekLabel(j+10))
		if j == 0 {
			for range 2 {
				weekDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, 0))
				week.WeekDays = append(week.WeekDays, *weekDay)
			}
		}
		block.Weeks = append(block.Weeks, *week)
	}

	utils.PrettyPrint(block)
	mockModels := []model.Block{*block}
	m.EXPECT().GetBlockWeeksByUserId("1").Return(mockModels, nil).Once()

	// Act
	actual, err := service.GetBlocks("1")
	utils.PrettyPrint(actual)

	// Assert
	require.Nil(t, err)
	assert.Equal(t, block.Id, actual[block.Label].Id)
	assert.Equal(t, block.Label, actual[block.Label].Label)
	assert.Equal(t, block.UserId, actual[block.Label].UserId)
	assert.Equal(t, block.Weeks[0], actual[block.Label].Weeks[10])
	assert.Equal(t, block.Weeks[1], actual[block.Label].Weeks[11])

}
