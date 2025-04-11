package exercise_type_handler

import (
	"trainer-helper/model"
)

type exericseTypeGetParams struct {
	UserId string `query:"user_id"`
}

type exerciseTypePostParams struct {
	Name         string           `json:"name"`
	Note         *string          `json:"note"`
	MediaType    *model.MediaType `json:"media_type"`
	MediaAddress *string          `json:"media_address"`
}

type exerciseTypePutPrams struct {
	Id           int              `json:"id"`
	Note         *string          `json:"note"`
	MediaType    *model.MediaType `json:"media_type"`
	MediaAddress *string          `json:"media_address"`
}

func (etpp exerciseTypePutPrams) toModel() model.ExerciseType {
	return model.ExerciseType{
		IdModel: model.IdModel{
			Id: etpp.Id,
		},
		Note:         etpp.Note,
		MediaType:    etpp.MediaType,
		MediaAddress: etpp.MediaAddress,
	}
}
