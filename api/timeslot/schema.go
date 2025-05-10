package timeslot

import (
	"fmt"
	"time"
	"trainer-helper/api"
	"trainer-helper/model"
)

type timeslotGetParams struct {
	StartDate time.Time `query:"start_date"`
	EndDate   time.Time `query:"end_date"`
}

type timeslotPostParams struct {
	TrainerId string    `json:"trainer_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func (tpp timeslotPostParams) ToModel() model.Timeslot {
	timeslotName := fmt.Sprintf("from %s to %s on %s",
		humanTime(tpp.Start),
		humanTime(tpp.End),
		humanDate(tpp.Start))
	return *model.BuildTimeslot(timeslotName, tpp.Start, tpp.End, tpp.TrainerId, nil)
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
		Name:      api.DerefString(tpp.Name),
		Start:     api.DerefTime(tpp.Start),
		End:       api.DerefTime(tpp.End),
	}
}

type timestlotUndeletePostParams struct {
	Id int `json:"id"`
}
