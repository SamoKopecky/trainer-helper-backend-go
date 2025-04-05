package exercise_duplicate_handler

type exerciseDuplicatePostParams struct {
	CopyTimeslotId int `json:"copy_timeslot_id"`
	TimeslotId     int `json:"timeslot_id"`
}
