package exercise_type

import (
	"trainer-helper/model"
)

type exericseTypeGetParams struct {
	UserId string `query:"user_id"`
}

type exerciseTypePostParams struct {
	Name        string           `json:"name"`
	Note        *string          `json:"note"`
	MediaType   *model.MediaType `json:"media_type"`
	YoutubeLink *string          `json:"youtube_link"`
	FilePath    *string          `json:"file_path"`
}

type exerciseTypePutPrams struct {
	Id          int              `json:"id"`
	Note        *string          `json:"note"`
	MediaType   *model.MediaType `json:"media_type"`
	YoutubeLink *string          `json:"youtube_link"`
	FilePath    *string          `json:"file_path"`
}

func (etpp exerciseTypePutPrams) ToModel() model.ExerciseType {
	return model.ExerciseType{
		IdModel: model.IdModel{
			Id: etpp.Id,
		},
		Note:        etpp.Note,
		MediaType:   etpp.MediaType,
		YoutubeLink: etpp.YoutubeLink,
		FilePath:    etpp.FilePath,
	}
}
