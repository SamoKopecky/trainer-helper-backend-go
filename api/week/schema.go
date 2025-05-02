package week

import (
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
)

type weekGetRequest struct {
	UserId string `query:"user_id"`
}

type weekPostRequest struct {
	UserId     string    `json:"user_id"`
	StartDate  time.Time `json:"start_date"`
	Label      int       `json:"label"`
	BlockLabel int       `json:"block_label"`
}

func (wpr weekPostRequest) toModel() *model.Week {
	return model.BuildWeek(wpr.UserId, wpr.StartDate, wpr.Label, wpr.BlockLabel, nil, nil, nil, nil, nil, nil, nil)
}

type weekDeleteRequest struct {
	Id int `json:"id"`
}

type weekPutRequest struct {
	Id        int        `json:"id"`
	StartDate *time.Time `json:"start_date"`
	Monday    *string    `json:"monday"`
	Tuesday   *string    `json:"tuesday"`
	Wednesday *string    `json:"wednesday"`
	Thursday  *string    `json:"thursday"`
	Friday    *string    `json:"friday"`
	Saturday  *string    `json:"saturday"`
	Sunday    *string    `json:"sunday"`
}

func (wpr weekPutRequest) toModel() *model.Week {
	return &model.Week{
		IdModel: model.IdModel{
			Id: wpr.Id,
		},
		StartDate: api.DerefTime(wpr.StartDate),
		Monday:    wpr.Monday,
		Tuesday:   wpr.Tuesday,
		Wednesday: wpr.Wednesday,
		Thursday:  wpr.Thursday,
		Friday:    wpr.Friday,
		Saturday:  wpr.Saturday,
		Sunday:    wpr.Sunday,
	}
}
