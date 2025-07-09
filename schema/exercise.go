package schema

type RawWorkSet struct {
	Reps      int    `json:"reps"`
	Intensity string `json:"intensity"`
	Rpe       *int   `json:"rpe"`
}

type RawExercise struct {
	Note         string       `json:"note"`
	ExerciseName string       `json:"exercise"`
	WorkSets     []RawWorkSet `json:"work_sets"`
}
