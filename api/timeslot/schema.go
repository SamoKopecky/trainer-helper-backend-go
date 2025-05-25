package timeslot

import (
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
	"trainer-helper/utils"
)

type timeslotEnchancedGetParams struct {
	Start time.Time `query:"start"`
	End   time.Time `query:"end"`
}

type timeslotGetParams struct {
	StartDate utils.Date `query:"start_date"`
	EndDate   utils.Date `query:"end_date"`
	UserId    string     `query:"user_id"`
}

type timeslotPostParams struct {
	TrainerId string    `json:"trainer_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func (tpp timeslotPostParams) ToModel() model.Timeslot {
	return *model.BuildTimeslot(tpp.Start, tpp.End, tpp.TrainerId, nil)
}

type timeslotDeleteParams struct {
	Id int `json:"id"`
}

type timeslotPutParams struct {
	Id        int        `json:"id"`
	TraineeId *string    `json:"trainee_id"`
	Name      *string    `json:"name"`
	Start     *time.Time `json:"start"`
	End       *time.Time `json:"end"`
}

func (tpp timeslotPutParams) ToModel() model.Timeslot {
	return model.Timeslot{
		IdModel: model.IdModel{
			Id: tpp.Id,
		},
		TraineeId: tpp.TraineeId,
		Start:     api.DerefTime(tpp.Start),
		End:       api.DerefTime(tpp.End),
	}
}

type timestlotUndeletePostParams struct {
	Id int `json:"id"`
}
