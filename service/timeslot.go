package service

import (
	"log"
	"time"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/schema"
	"trainer-helper/store"
)

type Timeslot struct {
	Crud    store.Timeslot
	Fetcher fetcher.IAM
}

func (t Timeslot) GetById(timeslotId int) (timeslot schema.Timeslot, err error) {
	crudTimeslot, err := t.Crud.GetById(timeslotId)
	if err != nil {
		return
	}
	timeslot = schema.Timeslot{Timeslot: crudTimeslot}
	if crudTimeslot.TraineeId == nil {
		return
	}

	user, err := t.Fetcher.GetUserById(*crudTimeslot.TraineeId)
	if err != nil {
		log.Fatal(err)
	}
	timeslot.UserName = user.FullName()
	timeslot.UserNickname = user.Nickname()
	return
}

func (t Timeslot) GetByRoleAndDate(start, end time.Time, users []fetcher.KeycloakUser, claims *schema.JwtClaims) ([]schema.Timeslot, error) {
	var err error
	var timeslots []model.Timeslot

	isTrainer := claims.IsTrainer()
	timeslots, err = t.Crud.GetByTimeRangeAndUserId(start, end, claims.Subject, isTrainer)
	if err != nil {
		return nil, err
	}

	iamUserMap := make(map[string]fetcher.KeycloakUser)
	for _, user := range users {
		iamUserMap[user.Id] = user
	}

	apiTimeslots := make([]schema.Timeslot, len(timeslots))
	for i, timeslot := range timeslots {
		apiTimeslot := schema.Timeslot{
			Timeslot: timeslot,
		}

		if apiTimeslot.TraineeId != nil {
			if user, ok := iamUserMap[*timeslot.TraineeId]; ok {
				apiTimeslot.UserName = user.FullName()
				apiTimeslot.UserNickname = user.Nickname()
			}
		}

		apiTimeslots[i] = apiTimeslot
	}

	return apiTimeslots, err

}
