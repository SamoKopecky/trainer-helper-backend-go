package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	bun.BaseModel `bun:"table:timeslot"`
	IdModel

	TrainerId string    `json:"trainer_id"`
	TraineeId *string   `json:"trainee_id"`
	Name      string    `json:"name"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Timestamp
	DeletedAt *time.Time `json:"-" bun:",soft_delete,nullzero"`
}

func BuildTimeslot(name string, start, end time.Time, deletedAt *time.Time, trainerId string, traineeId *string) *Timeslot {
	return &Timeslot{
		Name:      name,
		Start:     start,
		End:       end,
		TrainerId: trainerId,
		TraineeId: traineeId,
		Timestamp: buildTimestamp(),
		DeletedAt: deletedAt}
}

type ApiTimeslot struct {
	Timeslot
	UserName *string `json:"person_name"`
}
