package weekday

import (
	"time"
	"trainer-helper/model"
)

type weekDayPostRequest struct {
	WeekId  int       `json:"week_id"`
	UserId  string    `json:"user_id"`
	DayDate time.Time `json:"day_date"`
	Name    *string   `json:"name"`
}

func (wdpr weekDayPostRequest) ToModel() model.WeekDay {
	return *model.BuildWeekDay(wdpr.WeekId, wdpr.UserId, wdpr.DayDate, wdpr.Name)

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
