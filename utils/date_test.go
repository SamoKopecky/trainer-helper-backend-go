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
