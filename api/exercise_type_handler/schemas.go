package exercise_type_handler

import (
	"trainer-helper/model"
)

type exericseTypeGetParams struct {
	UserId string `query:"user_id"`
}

type exerciseTypePostParams struct {
	Name   string  `json:"name"`
	UserId string  `json:"userId"`
	Note   *string `json:"note"`
}

type exerciseTypePutPrams struct {
	Id   int     `json:"id"`
	Note *string `json:"note"`
}

func (etpp exerciseTypePutPrams) toModel() model.ExerciseType {
	return model.ExerciseType{
		IdModel: model.IdModel{
			Id: etpp.Id,
		},
		Note: etpp.Note,
	}
}
