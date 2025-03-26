package service

import (
	"log"
	"time"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/utils"
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

func (t Timeslot) GetByRoleAndDate(start, end time.Time, role string) ([]model.ApiTimeslot, error) {
	users, err := t.Fetcher.GetUsersByRole(role)
	if err != nil {
		return nil, err
	}
	utils.PrettyPrint(users)
	return nil, nil

}
