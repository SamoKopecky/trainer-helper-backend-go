package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type SetTypeEnum string

const (
	// Main compound exercises
	Squat      SetTypeEnum = "Squat"
	Deadlift   SetTypeEnum = "Deadlift"
	BenchPress SetTypeEnum = "Bench Press"

	// Variations and related exercises
	RomanianDeadlift SetTypeEnum = "RDL"
	HorizontalRow    SetTypeEnum = "Cable Horizontal Row"
	HackSquat        SetTypeEnum = "Hack Squat"
	LegPress         SetTypeEnum = "Leg Press"
	CalfRaise        SetTypeEnum = "Calf Raise"
	RingMuscleUp     SetTypeEnum = "Ring Muscle Up"
	PullUp           SetTypeEnum = "Pull Up"

	// Machine and isolation exercises
	MachineHipAbduction  SetTypeEnum = "Machine Hip Abduction"
	JeffersonCurl        SetTypeEnum = "Jefferson Curl"
	KettlebellSideBend   SetTypeEnum = "Kettlebell Side Bend"
	MachineChestPress    SetTypeEnum = "Machine Chest Press"
	Multipress           SetTypeEnum = "Multipress"
	Dips                 SetTypeEnum = "Dips"
	MachineShoulderPress SetTypeEnum = "Machine Shoulder Press"
	TricepsPushdown      SetTypeEnum = "Triceps Pushdown"
	BentArmLateralRaise  SetTypeEnum = "Bent Arm Lateral Raise"
	BenchCrunch          SetTypeEnum = "Bench Crunch"

	// General placeholder for unspecified exercises
	None SetTypeEnum = ""
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel

	TimeslotId int         `json:"timeslot_id"`
	GroupId    int         `json:"group_id"`
	Note       *string     `json:"note"`
	SetType    SetTypeEnum `json:"set_type"`
	Timestamp

	// Not used in DB model
	WorkSets []WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(timeslotId, groupId int, note string, setType SetTypeEnum) *Exercise {
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
