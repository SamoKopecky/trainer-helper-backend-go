package timeslot_handler

import (
	"time"
	"trainer-helper/model"
)

func humanTime(time time.Time) string {
	return time.Format("15:04")
}

func humanDate(time time.Time) string {
	return time.Format("02-01")
}

func toFullTimeslot(timeslot *model.Timeslot) (full model.ApiTimeslot, err error) {
	full = model.ApiTimeslot{
		Timeslot: *timeslot,
	}
	if timeslot.UserId == nil {
		return
	}

	if err != nil {
		return
	}
	return
}
