package crud

import (
	"strconv"
	"testing"
	"trainer-helper/model"
	"trainer-helper/testutil"

	"github.com/stretchr/testify/require"
)

func TestWeekGetByUserId(t *testing.T) {
	db := testSetupDb(t)
	crud := NewWeek(db)
	var weeks []model.Week
	for i := range 2 {
		week := testutil.WeekFactory(testutil.WeekUserId(strconv.Itoa(i)))

		crud.Insert(week)
		weeks = append(weeks, *week)
	}

	// Act
	dbModels, err := crud.GetByUserId(weeks[0].UserId)
	if err != nil {
		t.Fatalf("Failed to retrieve weeks: %v", err)
	}

	// Assert
	require.Len(t, dbModels, 1)

	assertWeek := weeks[0]
	assertWeek.SetZeroTimes()
	dbModels[0].SetZeroTimes()

	require.Equal(t, assertWeek, dbModels[0])
}
