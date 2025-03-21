package timeslot_handler

import (
	"time"
)

func humanTime(time time.Time) string {
	return time.Format("15:04")
}

func humanDate(time time.Time) string {
	return time.Format("02-01")
}

func derefString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func derefTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}
