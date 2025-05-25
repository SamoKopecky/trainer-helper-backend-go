package service

import (
	"time"
	"trainer-helper/fetcher"
	"trainer-helper/model"
	"trainer-helper/schema"
	"trainer-helper/store"
	"trainer-helper/utils"

	mapset "github.com/deckarep/golang-set/v2"
)

type Timeslot struct {
	TimeslotCrud store.Timeslot
	WeekDayCrud  store.WeekDay
	Fetcher      fetcher.IAM
}

func (t Timeslot) getWeekDayMap(timeslots []model.Timeslot) (map[int]model.WeekDay, error) {
	weekDayMap := make(map[int]model.WeekDay)

	timeslotIds := mapset.NewSet[int]()
	for _, timeslot := range timeslots {
		timeslotIds.Add(timeslot.Id)
	}

	weekDays, err := t.WeekDayCrud.GetByTimeslotIds(timeslotIds.ToSlice())
	if err != nil {
		return weekDayMap, err
	}

	for _, weekDay := range weekDays {
		if weekDay.TimeslotId != nil {
			weekDayMap[*weekDay.TimeslotId] = weekDay
		}
	}
	return weekDayMap, nil
}

func (t Timeslot) GetByRoleAndDate(start, end time.Time, users []fetcher.KeycloakUser, claims *schema.JwtClaims) ([]schema.Timeslot, error) {
	var err error
	var dbTimeslots []model.Timeslot

	isTrainer := claims.IsTrainer()
	dbTimeslots, err = t.TimeslotCrud.GetByTimeRangeAndUserId(start, end, claims.Subject, isTrainer)
	if err != nil {
		return nil, err
	}

	iamUserMap := make(map[string]fetcher.KeycloakUser)
	for _, user := range users {
		iamUserMap[user.Id] = user
	}
	weekDayMap, err := t.getWeekDayMap(dbTimeslots)
	if err != nil {
		return nil, err
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

		if weekDay, ok := weekDayMap[timeslot.Id]; ok {
			timeslot.WeekDay = &weekDay
		}

		timeslots[i] = timeslot
	}

	return timeslots, err

}

func (t Timeslot) GetByStartAndUser(start, end utils.Date, userId string) (timeslots []model.Timeslot, err error) {
	startDateTime, _ := start.ToTimerange()
	_, endDateTime := end.ToTimerange()
	timeslots, err = t.TimeslotCrud.GetByTimeRangeAndUserId(startDateTime, endDateTime, userId, false)

	return
}
