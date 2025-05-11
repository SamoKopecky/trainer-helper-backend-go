package week

import (
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
)

type weekPostRequest struct {
	BlockId int    `json:"block_id"`
	Label   int    `json:"label"`
	UserId  string `json:"user_id"`
	IsFirst bool   `json:"is_first"`
}

func (wpr weekPostRequest) ToModel() model.Week {
	return model.Week{
		BlockId: wpr.BlockId,
		Label:   wpr.Label,
		UserId:  wpr.UserId,
	}
}

type weekPutRequest struct {
	Id        int        `json:"id"`
	StartDate *time.Time `json:"start_date"`
}

func (wpr weekPutRequest) ToModel() model.Week {
	return model.Week{
		IdModel: model.IdModel{
			Id: wpr.Id,
		},
		StartDate: api.DerefTime(wpr.StartDate),
	}
}
