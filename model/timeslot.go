package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	bun.BaseModel `bun:"table:timeslot"`
	IdModel

	TrainerId int32     `json:"trainer_id"`
	UserId    *int32    `json:"user_id"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Timestamp
	DeletedAt *time.Time `json:"-" bun:",soft_delete,nullzero"`
}

func BuildTimeslot(name string, start, end time.Time, deletedAt *time.Time, trainerId int32, userId *int32) *Timeslot {
	return &Timeslot{
		Name:      name,
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		UserId:    userId,
		Timestamp: buildTimestamp(),
		DeletedAt: deletedAt}
}

type ApiTimeslot struct {
	Timeslot
	PersonName *string `json:"person_name"`
}
