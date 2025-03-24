package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type SetType string

const (
	Squat    SetType = "Squat"
	Deadlift SetType = "Deadlift"
	RDL      SetType = "RDL"
	None     SetType = ""
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel

	TimeslotId int32   `json:"timeslot_id"`
	GroupId    int32   `json:"group_id"`
	Note       *string `json:"note"`
	SetType    SetType `json:"set_type"`
	Timestamp

	// Not used in DB model
	WorkSets []*WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(timeslotId, groupId int32, note string, setType SetType) *Exercise {
	return &Exercise{
		TimeslotId: timeslotId,
		GroupId:    groupId,
		Note:       &note,
		SetType:    setType,
		Timestamp:  buildTimestamp(),
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
