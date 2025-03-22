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

func toFullTimeslot(timeslot *model.Timeslot, crudPerson crud.CRUDPerson) model.ApiTimeslot {
	full := model.ApiTimeslot{
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
