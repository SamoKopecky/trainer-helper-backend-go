package service

import (
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

func (t Timeslot) GetByRoleAndDate(start, end time.Time, users []fetcher.KeycloakUser, claims *schema.JwtClaims) ([]schema.Timeslot, error) {
	var err error
	var dbTimeslots []model.Timeslot

	isTrainer := claims.IsTrainer()
	dbTimeslots, err = t.Crud.GetByTimeRangeAndUserId(start, end, claims.Subject, isTrainer)
	if err != nil {
		return nil, err
	}

	iamUserMap := make(map[string]fetcher.KeycloakUser)
	for _, user := range users {
		iamUserMap[user.Id] = user
	}

	timeslots := make([]schema.Timeslot, len(dbTimeslots))
	for i, t := range dbTimeslots {
		timeslot := schema.Timeslot{
			Timeslot: t,
		}

		if timeslot.TraineeId != nil {
			if user, ok := iamUserMap[*t.TraineeId]; ok {
				userModel := user.ToUserModel()
				timeslot.User = &userModel
			}
		} else {
			timeslot.User = nil
		}

		timeslots[i] = timeslot
	}

	return timeslots, err

}
