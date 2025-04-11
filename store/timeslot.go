package store

import (
	"time"
	"trainer-helper/model"
)

type Timeslot interface {
	StoreBase[model.Timeslot]
	GetByTimeRangeAndUserId(startDate, endDate time.Time, trainerId string, isTrainer bool) ([]model.Timeslot, error)
	GetById(timeslotId int) (model.Timeslot, error)
	Delete(timeslotId int) error
	RevertSolfDelete(timeslotId int) error
}
