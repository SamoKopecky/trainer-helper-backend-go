package timeslot_handler

import (
	"time"
	"trainer-helper/model"
)

type timeslotGetParams struct {
	StartDate time.Time `query:"start_date"`
	EndDate   time.Time `query:"end_date"`
}

type timeslotPostParams struct {
	TrainerId int32     `json:"trainer_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

type timeslotDeleteParams struct {
	Id int32 `json:"id"`
}

type timeslotPutParams struct {
	Id     int32      `json:"id"`
	UserId *int32     `json:"user_id"`
	Name   *string    `json:"name"`
	Start  *time.Time `json:"start"`
	End    *time.Time `json:"end"`
}

func (tpp timeslotPutParams) toModel() model.Timeslot {
	return model.Timeslot{
		IdModel: model.IdModel{
			Id: tpp.Id,
		},
		UserId: tpp.UserId,
		Name:   derefString(tpp.Name),
		Start:  derefTime(tpp.Start),
		End:    derefTime(tpp.End),
	}
}
