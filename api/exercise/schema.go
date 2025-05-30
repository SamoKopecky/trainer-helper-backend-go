package exercise

import (
	"trainer-helper/api"
	"trainer-helper/model"
)

type exerciseGetParams struct {
	WeekDayIds []int `query:"week_day_ids"`
}

type exercisePutParams struct {
	Id             int     `json:"id"`
	GroupId        *int    `json:"group_id"`
	ExerciseTypeId *int    `json:"exercise_type_id"`
	Note           *string `json:"note"`
}

type exerciseDuplicatePostParams struct {
	CopyTimeslotId int `json:"copy_timeslot_id"`
	TimeslotId     int `json:"timeslot_id"`
}

type exerciseCountPostParams struct {
	Count           int           `json:"count"`
	WorkSetTemplate model.WorkSet `json:"work_set_template"`
}

type exerciseCountDeleteParams struct {
	WorkSetIds []int `json:"work_set_ids"`
}

type exercisePostParams struct {
	WeekDayId int `json:"week_day_id"`
	GroupId   int `json:"group_id"`
}

func (epp exercisePutParams) ToModel() model.Exercise {
	return model.Exercise{
		IdModel: model.IdModel{
			Id: epp.Id,
		},
		GroupId:        api.DerefInt(epp.GroupId),
		ExerciseTypeId: epp.ExerciseTypeId,
		Note:           epp.Note,
	}
}

type exerciseUndeletePostParams struct {
	Id int `json:"id"`
}
