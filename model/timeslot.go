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
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func BuildTimeslot(name string, start, end time.Time, trainerId string, traineeId *string) *Timeslot {
	return &Timeslot{
		Name:      name,
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		TraineeId: traineeId,
		Timestamp: buildTimestamp(),
	}
}
