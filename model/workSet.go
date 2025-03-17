package model

import (
	"github.com/uptrace/bun"
)

type WorkSet struct {
	bun.BaseModel `bun:"table:work_set"`

	Id         int32 `bun:",pk,autoincrement"`
	ExerciseId int32
	reps       int32
	intensity  string
	rpe        int32
	Timestamp
}

func BuildWorkset(exerciseId, reps, rpe int32, intensity string) *WorkSet {
	return &WorkSet{
		ExerciseId: exerciseId,
		reps:       reps,
		rpe:        rpe,
		intensity:  intensity,
		Timestamp:  buildTimestamp()}
}
