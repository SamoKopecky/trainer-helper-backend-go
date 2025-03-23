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

type ExerciseWorkSets struct {
	Exercise
	WorkSetCount int32     `json:"work_set_count"`
	WorkSets     []WorkSet `json:"work_sets"`
}

type TimeslotExercises struct {
	Timeslot  ApiTimeslot         `json:"timeslot"`
	Exercises []*ExerciseWorkSets `json:"exercises"`
}

type CRUDExerciseWorkSets struct {
	ExerciseId int32
	WorkSetId  int32
	Exercise
	WorkSet
}

func (ews ExerciseWorkSets) SortWorkSets() {
	sort.Slice(ews.WorkSets, func(i, j int) bool {
		return ews.WorkSets[i].Id < ews.WorkSets[j].Id
	})
}

func (cews CRUDExerciseWorkSets) ToWorkSet() WorkSet {
	res := cews.WorkSet
	res.Id = cews.WorkSetId
	res.ExerciseId = cews.ExerciseId
	return res
}

func (cews CRUDExerciseWorkSets) ToExercise() Exercise {
	res := cews.Exercise
	res.Id = cews.ExerciseId
	return res
}
