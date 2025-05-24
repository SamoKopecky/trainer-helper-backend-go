package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	bun.BaseModel `bun:"table:timeslot"`
	IdModel
	Timestamp
	DeletedTimestamp

	TrainerId string    `json:"trainer_id"`
	TraineeId *string   `json:"trainee_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func BuildTimeslot(start, end time.Time, trainerId string, traineeId *string) *Timeslot {
	return &Timeslot{
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		TraineeId: traineeId,
		Timestamp: buildTimestamp(),
	}
}
