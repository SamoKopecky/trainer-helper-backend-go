package week

import (
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/utils"
)

type WeekDuplicatePostRequest struct {
	TemplateWeekId int `json:"template_week_id"`
	NewWeekId      int `json:"new_week_id"`
}

type WeekGetRequest struct {
	UserId    string     `query:"user_id"`
	StartDate utils.Date `query:"start_date"`
}

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
	Id        int         `json:"id"`
	StartDate *utils.Date `json:"start_date"`
}

func (wpr weekPutRequest) ToModel() model.Week {
	return model.Week{
		IdModel: model.IdModel{
			Id: wpr.Id,
		},
		StartDate: api.DerefDate(wpr.StartDate),
	}
}
