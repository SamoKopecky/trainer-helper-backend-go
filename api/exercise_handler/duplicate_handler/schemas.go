package exercise_duplicate_handler

type exerciseDuplicatePostParams struct {
	CopyTimeslotId int32 `json:"copy_timeslot_id"`
	TimeslotId     int32 `json:"timeslot_id"`
}
