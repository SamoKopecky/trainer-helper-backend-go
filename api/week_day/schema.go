package weekday

import (
	"trainer-helper/model"
	"trainer-helper/utils"
)

type weekDayGetRequest struct {
	WeekId int `query:"week_id"`
}

type weekDayPostRequest struct {
	WeekId  int        `json:"week_id"`
	UserId  string     `json:"user_id"`
	DayDate utils.Date `json:"day_date"`
	Name    *string    `json:"name"`
}

func (wdpr weekDayPostRequest) ToModel() model.WeekDay {
	return *model.BuildWeekDay(wdpr.WeekId, wdpr.UserId, wdpr.DayDate.Time, wdpr.Name)

}

type weekDayPutRequest struct {
	Id   int     `json:"id"`
	Name *string `json:"name"`
}

func (wdpr weekDayPutRequest) ToModel() model.WeekDay {
	return model.WeekDay{
		IdModel: model.IdModel{
			Id: wdpr.Id,
		},
		Name: wdpr.Name,
	}
}
