package model

import (
	"github.com/uptrace/bun"
)

type SetType string

const (
	Squat    SetType = "Squat"
	Deadlift SetType = "Deadlift"
	RDL      SetType = "RDL"
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`

	Id          int32 `bun:",pk,autoincrement"`
	TimestlotId int32
	GroupId     int32
	note        *string
	SetType
	Timestamp
}

func BuildExercise(timeslotId, groupId int32, note string, setType SetType) *Exercise {
	return &Exercise{
		TimestlotId: timeslotId,
		GroupId:     groupId,
		note:        &note,
		SetType:     setType,
		Timestamp:   buildTimestamp(),
	}
}
