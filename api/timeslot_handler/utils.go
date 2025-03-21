package timeslot_handler

import (
	"log"
	"time"
	"trainer-helper/crud"
	"trainer-helper/model"
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

func toFullTimeslot(timeslot *model.Timeslot, crudPerson crud.CRUDPerson) model.TimeslotFull {
	full := model.TimeslotFull{
		Timeslot: *timeslot,
	}
	if timeslot.UserId == nil {
		return full
	}

	person, err := crudPerson.Get(*timeslot.UserId)
	if err != nil {
		log.Fatal(err)
	}
	full.PersonName = &person.Name
	return full
}
