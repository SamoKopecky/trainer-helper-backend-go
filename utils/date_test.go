package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextMonday(t *testing.T) {
	testCases := []struct {
		date     time.Time
		expected time.Time
		name     string
	}{
		{
			name:     "saturday start",
			date:     time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 5, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "monday start",
			date:     time.Date(2025, 5, 12, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 5, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "tuesday start",
			date:     time.Date(2025, 5, 13, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 5, 19, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetNextMonday(tc.date)
			assert.Equal(t, tc.expected.Day(), actual.Day())
		})
	}
}

func TestDateTimerange(t *testing.T) {
	dateTime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	date := Date{Time: dateTime}
	start, end := date.ToTimerange()

	assert.Equal(t, "2006-01-02 00:00:00", start.Format("2006-01-02 15:04:05"))
	assert.Equal(t, "2006-01-02 23:59:59", end.Format("2006-01-02 15:04:05"))
}
