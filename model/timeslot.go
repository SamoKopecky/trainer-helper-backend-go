package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	bun.BaseModel `bun:"table:timeslot"`

	Id        int32 `bun:",pk,autoincrement"`
	TrainerId int32
	UserId    int32
	Name      string
	Start     time.Time
	End       time.Time
	Timestamp
}

func BuildTimeslot(name, email string, start, end time.Time, trainerId, userId int32) *Timeslot {
	return &Timeslot{
		Name:      name,
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		UserId:    userId,
		Timestamp: buildTimestamp()}
}
