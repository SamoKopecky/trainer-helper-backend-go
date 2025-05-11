package utils

import "time"

func GetNextMonday(t time.Time) time.Time {
	currentWeekday := t.Weekday()
	daysToAdd := int(time.Monday - currentWeekday)
	if daysToAdd < 0 {
		daysToAdd += 7
	}

	return t.AddDate(0, 0, daysToAdd)
}
