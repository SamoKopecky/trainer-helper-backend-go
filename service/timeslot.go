package service

import (
	"log"
	"trainer-helper/crud"
	"trainer-helper/fetcher"
	"trainer-helper/model"
)

type Timeslot struct {
	Crud crud.CRUDTimeslot
	IAM  fetcher.IAM
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

	iamTimeslot, err := t.IAM.GetUserById(*crudTimeslot.TraineeId)
	if err != nil {
		log.Fatal(err)
	}
	fullName := iamTimeslot.FullName()
	timeslot.UserName = &fullName
	return
}
