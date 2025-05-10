package weekday

import (
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
)

type weekDayPostRequest struct {
	WeekId  int       `json:"week_id"`
	UserId  string    `json:"user_id"`
	Name    string    `json:"name"`
	DayDate time.Time `json:"day_date"`
}

func (wdpr weekDayPostRequest) ToModel() model.WeekDay {
	return *model.BuildWeekDay(wdpr.WeekId, wdpr.UserId, wdpr.Name, wdpr.DayDate)

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
		Name: api.DerefString(wdpr.Name),
	}
}
