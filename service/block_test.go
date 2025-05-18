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
	blockWithWeeks := testutil.BlockFactory(testutil.BlockUserId(userId), testutil.BlockLabel(3))
	blockWithNoWeeks := testutil.BlockFactory(testutil.BlockUserId(userId), testutil.BlockLabel(2))

	weekOne := testutil.WeekFactory(testutil.WeekIds(userId, 0), testutil.WeekLabel(30))
	weekTwo := testutil.WeekFactory(testutil.WeekIds(userId, 0), testutil.WeekLabel(20))

	blockWithWeeks.Weeks = []model.Week{*weekOne, *weekTwo}

	mockModels := []model.Block{*blockWithWeeks, *blockWithNoWeeks}
	m.EXPECT().GetBlockWeeksByUserId("1").Return(mockModels, nil).Once()

	// Act
	actual, err := service.GetBlocks("1")

	// Assert

	require.Nil(t, err)
	assert.Equal(t, 2, actual[0].Label)
	assert.Equal(t, 0, len(actual[0].Weeks))
	assert.Equal(t, 3, actual[1].Label)
	assert.Equal(t, 2, len(actual[1].Weeks))
	assert.Equal(t, 20, actual[1].Weeks[0].Label)
	assert.Equal(t, 0, len(actual[1].Weeks[0].WeekDays))
	assert.Equal(t, 30, actual[1].Weeks[1].Label)
	assert.Equal(t, 0, len(actual[1].Weeks[1].WeekDays))
}
