package schema

import "trainer-helper/model"

type RawWorkSet struct {
	Reps      int     `json:"reps"`
	Intensity string  `json:"intensity"`
	Rpe       *string `json:"rpe"`
}

type RawExercise struct {
	Note         string       `json:"note"`
	ExerciseName string       `json:"exercise_name"`
	WorkSets     []RawWorkSet `json:"work_sets"`
}

func (rws RawWorkSet) ToModel(exercseId int) model.WorkSet {
	return model.WorkSet{
		Intensity:  rws.Intensity,
		ExerciseId: exercseId,
		Reps:       rws.Reps,
		Rpe:        rws.Rpe,
	}
}

func (re RawExercise) ToModel(weekDayId, groupId int) model.Exercise {
	return model.Exercise{
		WeekDayId: weekDayId,
		Note:      &re.Note,
		GroupId:   groupId,
	}
}
