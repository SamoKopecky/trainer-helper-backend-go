package exercise_handler

import (
	"trainer-helper/api"
	"trainer-helper/model"
)

type exerciseGetParams struct {
	Id int
}

type exercisePutParams struct {
	Id      int                `json:"id"`
	GroupId *int               `json:"group_id"`
	ExerciseType *model.ExerciseTypeEnum `json:"exercise_type"`
	Note    *string            `json:"note"`
}

type exerciseDeleteParams struct {
	TimeslotId int `json:"timeslot_id"`
	ExerciseId int `json:"exercise_id"`
}

type exercisePostParams struct {
	TimeslotId int `json:"timeslot_id"`
	GroupId    int `json:"group_id"`
}

func (epp exercisePutParams) toModel() model.Exercise {
	return model.Exercise{
		IdModel: model.IdModel{
			Id: epp.Id,
		},
		GroupId: api.DerefInt(epp.GroupId),
		ExerciseType: api.DerefExerciseType(epp.ExerciseType),
		Note:    epp.Note,
	}
}
