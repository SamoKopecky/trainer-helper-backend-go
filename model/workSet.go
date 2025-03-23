package model

import (
	"github.com/uptrace/bun"
)

type WorkSet struct {
	bun.BaseModel `bun:"table:work_set"`
	IdModel

	ExerciseId int32  `json:"exercise_id"`
	Reps       int32  `json:"reps"`
	Intensity  string `json:"intensity"`
	Rpe        *int32 `json:"rpe"`
	Timestamp
}

func BuildWorkSet(exerciseId, reps int32, rpe *int32, intensity string) *WorkSet {
	return &WorkSet{
		ExerciseId: exerciseId,
		Reps:       reps,
		Rpe:        rpe,
		Intensity:  intensity,
		Timestamp:  buildTimestamp()}
}
