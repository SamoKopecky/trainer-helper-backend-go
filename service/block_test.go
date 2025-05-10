package service

import (
	"testing"
	"time"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/testutil"
	"trainer-helper/utils"

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

	now := time.Now()
	tommorow := now.Add(time.Hour * 24)
	dayAfterTommorow := now.Add(time.Hour * 48)
	weekOneDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, 0), testutil.WeekDayTime(tommorow))
	weekOne.WeekDays = append(weekOne.WeekDays, *weekOneDay)

	weekThreeDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, 0), testutil.WeekDayTime(dayAfterTommorow))
	weekOne.WeekDays = append(weekOne.WeekDays, *weekThreeDay)

	weekTwoDay := testutil.WeekDayFactory(testutil.WeekDayIds(userId, 0), testutil.WeekDayTime(now))
	weekOne.WeekDays = append(weekOne.WeekDays, *weekTwoDay)

	blockWithWeeks.Weeks = []model.Week{*weekOne, *weekTwo}

	mockModels := []model.Block{*blockWithWeeks, *blockWithNoWeeks}
	m.EXPECT().GetBlockWeeksByUserId("1").Return(mockModels, nil).Once()

	// Act
	actual, err := service.GetBlocks("1")
	utils.PrettyPrint(actual)

	// Assert

	require.Nil(t, err)
	assert.Equal(t, 2, actual[0].Label)
	assert.Equal(t, 0, len(actual[0].Weeks))
	assert.Equal(t, 3, actual[1].Label)
	assert.Equal(t, 2, len(actual[1].Weeks))
	assert.Equal(t, 20, actual[1].Weeks[0].Label)
	assert.Equal(t, 0, len(actual[1].Weeks[0].WeekDays))
	assert.Equal(t, 30, actual[1].Weeks[1].Label)
	assert.Equal(t, 3, len(actual[1].Weeks[1].WeekDays))
	assert.Equal(t, now, actual[1].Weeks[1].WeekDays[0].DayDate)
	assert.Equal(t, tommorow, actual[1].Weeks[1].WeekDays[1].DayDate)
	assert.Equal(t, dayAfterTommorow, actual[1].Weeks[1].WeekDays[2].DayDate)
}
