package weekday

import (
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/utils"
)

type weekDayPostFromRawRequest struct {
	RawData string `json:"raw_data"`
}

type weekDayGetRequest struct {
	WeekId *int `query:"week_id"`

	DayDate *utils.Date `query:"day_date"`
	UserId  *string     `query:"user_id"`
}

type weekDayPostRequest struct {
	WeekId  int        `json:"week_id"`
	UserId  string     `json:"user_id"`
	DayDate utils.Date `json:"day_date"`
	Name    *string    `json:"name"`
}

func (wdpr weekDayPostRequest) ToModel() model.WeekDay {
	return *model.BuildWeekDay(wdpr.WeekId, wdpr.UserId, wdpr.DayDate.Time, wdpr.Name, nil)

}

type weekDayPutRequest struct {
	Id         int         `json:"id"`
	Name       *string     `json:"name"`
	DayDate    *utils.Date `json:"day_date"`
	TimeslotId *int        `json:"timeslot_id"`
}

func (wdpr weekDayPutRequest) ToModel() model.WeekDay {
	return model.WeekDay{
		IdModel: model.IdModel{
			Id: wdpr.Id,
		},
		Name:       wdpr.Name,
		TimeslotId: wdpr.TimeslotId,
		DayDate:    api.DerefDate(wdpr.DayDate),
	}
}
