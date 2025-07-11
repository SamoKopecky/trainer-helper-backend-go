package work_set

import (
	"trainer-helper/api"
	"trainer-helper/model"
)

type workSetPutRequest struct {
	Id       int     `json:"id"`
	Reps     *int    `json:"reps"`
	Inensity *string `json:"intensity"`
	Rpe      *string `json:"rpe"`
}

func (wspr workSetPutRequest) ToModel() model.WorkSet {
	return model.WorkSet{
		IdModel: model.IdModel{
			Id: wspr.Id,
		},
		Rpe:       wspr.Rpe,
		Reps:      api.DerefInt(wspr.Reps),
		Intensity: api.DerefString(wspr.Inensity),
	}
}

type workSetUndeletePostParams struct {
	Ids []int `json:"ids"`
}
