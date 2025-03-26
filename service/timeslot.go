package service

import (
	"log"
	"time"
	"trainer-helper/api"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/model"
)

type Timeslot struct {
	Crud    crud.Timeslot
	Fetcher fetcher.IAM
}

func (t Timeslot) GetById(timeslotId int32) (timeslot model.ApiTimeslot, err error) {
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

func (t Timeslot) GetByRoleAndDate(start, end time.Time, claims *api.JwtClaims) ([]model.ApiTimeslot, error) {
	var err error
	var iamUsers []fetcher.KeycloakUser
	var timeslots []model.Timeslot

	role, isTrainer := claims.GetAppRole()
	if isTrainer {
		iamUsers, err = t.Fetcher.GetUsersByRole(role)
		if err != nil {
			return nil, err
		}
		timeslots, err = t.Crud.GetByTimeRangeAndTrainerId(start, end, claims.Subject)
		if err != nil {
			return nil, err
		}
	} else {
		user, err := t.Fetcher.GetUserById(claims.Subject)
		if err != nil {
			return nil, err
		}
		iamUsers = append(iamUsers, user)
		timeslots, err = t.Crud.GetByTimeRangeAndTraineeId(start, end, claims.Subject)
		if err != nil {
			return nil, err
		}

	}

	iamUserMap := make(map[string]fetcher.KeycloakUser)
	for _, user := range iamUsers {
		iamUserMap[user.Id] = user
	}

	apiTimeslots := make([]model.ApiTimeslot, len(timeslots))
	for i, timeslot := range timeslots {
		apiTimeslot := model.ApiTimeslot{
			Timeslot: timeslot,
		}

		if user, ok := iamUserMap[*timeslot.TraineeId]; ok {
			fullName := user.FullName()
			apiTimeslot.UserName = &fullName
		}

		apiTimeslots[i] = apiTimeslot
	}

	return apiTimeslots, err

}
