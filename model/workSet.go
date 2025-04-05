package model

import (
	"github.com/uptrace/bun"
)

type WorkSet struct {
	bun.BaseModel `bun:"table:work_set"`
	IdModel

	ExerciseId int    `json:"exercise_id"`
	Reps       int    `json:"reps"`
	Intensity  string `json:"intensity"`
	Rpe        *int   `json:"rpe"`
	Timestamp
}

func BuildWorkSet(exerciseId, reps int, rpe *int, intensity string) *WorkSet {
	return &WorkSet{
		ExerciseId: exerciseId,
		Reps:       reps,
		Rpe:        rpe,
		Intensity:  intensity,
		Timestamp:  buildTimestamp()}
}

func (ws *WorkSet) ToNew(exerciseId int) {
	ws.Id = EmptyId
	ws.Timestamp = buildTimestamp()
	ws.ExerciseId = exerciseId
}
