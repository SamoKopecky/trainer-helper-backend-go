package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	TimeslotId     int     `json:"timeslot_id"`
	GroupId        int     `json:"group_id"`
	Note           *string `json:"note"`
	ExerciseTypeId *int    `json:"exercise_type_id"`

	// Not used in DB model
	WorkSets []WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(timeslotId, groupId int, note *string, exerciseTypeId *int) *Exercise {
	return &Exercise{
		TimeslotId:     timeslotId,
		GroupId:        groupId,
		Note:           note,
		ExerciseTypeId: exerciseTypeId,
		Timestamp:      buildTimestamp(),
	}
}

type TimeslotExercises struct {
	Timeslot  ApiTimeslot `json:"timeslot"`
	Exercises []*Exercise `json:"exercises"`
}

func (e Exercise) SortWorkSets() {
	sort.Slice(e.WorkSets, func(i, j int) bool {
		return e.WorkSets[i].Id < e.WorkSets[j].Id
	})
}

func (e *Exercise) ToNew(timeslotId int) {
	e.Id = EmptyId
	e.Timestamp = buildTimestamp()
	e.TimeslotId = timeslotId
}
