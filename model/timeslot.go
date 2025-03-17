package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	bun.BaseModel `bun:"table:timeslot"`

	Id        int32     `bun:",pk,autoincrement" json:"id"`
	TrainerId int32     `json:"trainer_id"`
	UserId    *int32    `json:"user_id"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Timestamp
}

func BuildTimeslot(name string, start, end time.Time, trainerId int32, userId *int32) *Timeslot {
	return &Timeslot{
		Name:      name,
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		UserId:    userId,
		Timestamp: buildTimestamp()}
}
