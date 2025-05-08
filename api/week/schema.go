package week

import (
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
)

type weekPostRequest struct {
	BlockId   int       `json:"block_id"`
	StartDate time.Time `json:"start_date"`
	Label     int       `json:"label"`
	UserId    string    `json:"user_id"`
}

func (wpr weekPostRequest) toModel() *model.Week {
	return model.BuildWeek(wpr.BlockId, wpr.StartDate, wpr.Label, wpr.UserId)
}

type weekDeleteRequest struct {
	Id int `json:"id"`
}

type weekPutRequest struct {
	Id        int        `json:"id"`
	StartDate *time.Time `json:"start_date"`
}

func (wpr weekPutRequest) toModel() *model.Week {
	return &model.Week{
		IdModel: model.IdModel{
			Id: wpr.Id,
		},
		StartDate: api.DerefTime(wpr.StartDate),
	}
}
