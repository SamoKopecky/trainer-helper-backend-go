package timeslot

import "time"

func humanTime(time time.Time) string {
	return time.Format("15:04")
}

func humanDate(time time.Time) string {
	return time.Format("02-01")
}
