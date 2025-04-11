package service

import (
	"log"
	"time"
	"trainer-helper/api"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/store"
)

type Timeslot struct {
	Crud    store.Timeslot
	Fetcher fetcher.IAM
}

func (t Timeslot) GetById(timeslotId int) (timeslot model.ApiTimeslot, err error) {
	crudTimeslot, err := t.Crud.GetById(timeslotId)
	if err != nil {
		return
	}
	timeslot = model.ApiTimeslot{Timeslot: crudTimeslot}
	if crudTimeslot.TraineeId == nil {
		return
	}

	iamTimeslot, err := t.Fetcher.GetUserById(*crudTimeslot.TraineeId)
	if err != nil {
		log.Fatal(err)
	}
	fullName := iamTimeslot.FullName()
	timeslot.UserName = &fullName
	return
}

func (t Timeslot) GetByRoleAndDate(start, end time.Time, users []fetcher.KeycloakUser, claims *api.JwtClaims) ([]model.ApiTimeslot, error) {
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

	apiTimeslots := make([]model.ApiTimeslot, len(timeslots))
	for i, timeslot := range timeslots {
		apiTimeslot := model.ApiTimeslot{
			Timeslot: timeslot,
		}

		if apiTimeslot.TraineeId != nil {
			if user, ok := iamUserMap[*timeslot.TraineeId]; ok {
				fullName := user.FullName()
				apiTimeslot.UserName = &fullName
			}
		}

		apiTimeslots[i] = apiTimeslot
	}

	return apiTimeslots, err

}
