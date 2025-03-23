package exercise_handler

import (
	"trainer-helper/api"
	"trainer-helper/model"
)

type exerciseGetParams struct {
	Id int32
}

type exercisePutParams struct {
	Id      int32          `json:"id"`
	GroupId *int32         `json:"group_id"`
	SetType *model.SetType `json:"set_type"`
	Note    *string        `json:"note"`
}

type exerciseDeleteParams struct {
	TimeslotId int32 `json:"timeslot_id"`
	ExerciseId int32 `json:"exercise_id"`
}

func (epp exercisePutParams) toModel() model.Exercise {
	return model.Exercise{
		IdModel: model.IdModel{
			Id: epp.Id,
		},
		GroupId: api.DerefInt(epp.GroupId),
		SetType: api.DerefSetType(epp.SetType),
		Note:    epp.Note,
	}
}
