package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp
	DeletedTimestamp

	WeekDayId      int     `json:"week_day_id"`
	GroupId        int     `json:"group_id"`
	Note           *string `json:"note"`
	ExerciseTypeId *int    `json:"exercise_type_id"`

	// Not used in DB model
	WorkSets []WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(weekDayId, groupId int, note *string, exerciseTypeId *int) *Exercise {
	return &Exercise{
		WeekDayId:      weekDayId,
		GroupId:        groupId,
		Note:           note,
		ExerciseTypeId: exerciseTypeId,
		Timestamp:      buildTimestamp(),
	}
}

func (e Exercise) SortWorkSets() {
	if e.WorkSets == nil {
		return
	}
	sort.Slice(e.WorkSets, func(i, j int) bool {
		return e.WorkSets[i].Id < e.WorkSets[j].Id
	})
}

func (e *Exercise) ToNew(weekDayId int) {
	e.Id = EmptyId
	e.Timestamp = buildTimestamp()
	e.WeekDayId = weekDayId
}
