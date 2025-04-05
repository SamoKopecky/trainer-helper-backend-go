package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type SetType string

const (
	// Main compound exercises
	Squat      SetType = "Squat"
	Deadlift   SetType = "Deadlift"
	BenchPress SetType = "Bench Press"

	// Variations and related exercises
	RomanianDeadlift SetType = "RDL"
	HorizontalRow    SetType = "Cable Horizontal Row"
	HackSquat        SetType = "Hack Squat"
	LegPress         SetType = "Leg Press"
	CalfRaise        SetType = "Calf Raise"
	RingMuscleUp     SetType = "Ring Muscle Up"
	PullUp           SetType = "Pull Up"

	// Machine and isolation exercises
	MachineHipAbduction  SetType = "Machine Hip Abduction"
	JeffersonCurl        SetType = "Jefferson Curl"
	KettlebellSideBend   SetType = "Kettlebell Side Bend"
	MachineChestPress    SetType = "Machine Chest Press"
	Multipress           SetType = "Multipress"
	Dips                 SetType = "Dips"
	MachineShoulderPress SetType = "Machine Shoulder Press"
	TricepsPushdown      SetType = "Triceps Pushdown"
	BentArmLateralRaise  SetType = "Bent Arm Lateral Raise"
	BenchCrunch          SetType = "Bench Crunch"

	// General placeholder for unspecified exercises
	None SetType = ""
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel

	TimeslotId int     `json:"timeslot_id"`
	GroupId    int     `json:"group_id"`
	Note       *string `json:"note"`
	SetType    SetType `json:"set_type"`
	Timestamp

	// Not used in DB model
	WorkSets []WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(timeslotId, groupId int, note string, setType SetType) *Exercise {
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

func (e *Exercise) ToNew(timeslotId int) {
	e.Id = EmptyId
	e.Timestamp = buildTimestamp()
	e.TimeslotId = timeslotId
}
