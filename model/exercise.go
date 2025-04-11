package model

import (
	"sort"

	"github.com/uptrace/bun"
)

type ExerciseTypeEnum string

const (
	// Main compound exercises
	Squat      ExerciseTypeEnum = "Squat"
	Deadlift   ExerciseTypeEnum = "Deadlift"
	BenchPress ExerciseTypeEnum = "Bench Press"

	// Variations and related exercises
	RomanianDeadlift ExerciseTypeEnum = "RDL"
	HorizontalRow    ExerciseTypeEnum = "Cable Horizontal Row"
	HackSquat        ExerciseTypeEnum = "Hack Squat"
	LegPress         ExerciseTypeEnum = "Leg Press"
	CalfRaise        ExerciseTypeEnum = "Calf Raise"
	RingMuscleUp     ExerciseTypeEnum = "Ring Muscle Up"
	PullUp           ExerciseTypeEnum = "Pull Up"

	// Machine and isolation exercises
	MachineHipAbduction  ExerciseTypeEnum = "Machine Hip Abduction"
	JeffersonCurl        ExerciseTypeEnum = "Jefferson Curl"
	KettlebellSideBend   ExerciseTypeEnum = "Kettlebell Side Bend"
	MachineChestPress    ExerciseTypeEnum = "Machine Chest Press"
	Multipress           ExerciseTypeEnum = "Multipress"
	Dips                 ExerciseTypeEnum = "Dips"
	MachineShoulderPress ExerciseTypeEnum = "Machine Shoulder Press"
	TricepsPushdown      ExerciseTypeEnum = "Triceps Pushdown"
	BentArmLateralRaise  ExerciseTypeEnum = "Bent Arm Lateral Raise"
	BenchCrunch          ExerciseTypeEnum = "Bench Crunch"

	// General placeholder for unspecified exercises
	None ExerciseTypeEnum = ""
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel

	TimeslotId int         `json:"timeslot_id"`
	GroupId    int         `json:"group_id"`
	Note       *string     `json:"note"`
	ExerciseType    ExerciseTypeEnum `json:"exercise_type"`
	Timestamp

	// Not used in DB model
	WorkSets []WorkSet `bun:"rel:has-many,join:id=exercise_id" json:"work_sets"`
}

func BuildExercise(timeslotId, groupId int, note string, ExerciseType ExerciseTypeEnum) *Exercise {
	return &Exercise{
		TimeslotId: timeslotId,
		GroupId:    groupId,
		Note:       &note,
		ExerciseType:    ExerciseType,
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
