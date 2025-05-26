package service

import (
	"time"
	"trainer-helper/model"
	"trainer-helper/store"
	"trainer-helper/utils"
)

const DaysInAWeek = 7

type Week struct {
	WeekStore     store.Week
	WeekDayStore  store.WeekDay
	ExerciseStore store.Exercise
	WorkSetStore  store.WorkSet
}

func (w Week) CreateWeek(newWeek *model.Week, isFirst bool) (err error) {
	var lastDate time.Time
	if isFirst {
		lastDate, err = w.WeekStore.GetPreviousBlockId(newWeek.UserId)
	} else {
		lastDate, err = w.WeekStore.GetLastWeekDate(newWeek.BlockId)
	}
	if err != nil {
		return err
	}

	if lastDate.IsZero() {
		// If no week exists, just get next monday
		newWeek.StartDate = utils.GetNextMonday(time.Now())
	} else {
		// If some week already exists, add 7 days to get next monday
		newWeek.StartDate = utils.GetNextMonday(lastDate).AddDate(0, 0, 7)
	}

	if err = w.WeekStore.Insert(newWeek); err != nil {
		return
	}
	newWeek.WeekDays = []model.WeekDay{}

	return nil

}

func (w Week) DuplicateWeekDays(templateWeekId, newWeekId int) error {
	weekDays, templateWeekDays, err := w.duplicateWeekDays(templateWeekId, newWeekId)
	if err != nil {
		return err
	}
	templateWeekDayIds := make([]int, len(weekDays))
	for i, weekDay := range templateWeekDays {
		templateWeekDayIds[i] = weekDay.Id
	}

	templateExercises, err := w.ExerciseStore.GetExerciseWorkSets(templateWeekDayIds)
	templateExercisesMap := make(map[int][]model.Exercise)
	for _, e := range templateExercises {
		templateExercisesMap[e.WeekDayId] = append(templateExercisesMap[e.WeekDayId], e)
	}

	var workSets []model.WorkSet
	for i, weekDay := range weekDays {
		wdExercises := templateExercisesMap[templateWeekDays[i].Id]
		for _, exercise := range wdExercises {
			exercise.Id = 0
			exercise.SetZeroTimes()
			exercise.WeekDayId = weekDay.Id
			err := w.ExerciseStore.Insert(&exercise)
			if err != nil {
				return err
			}

			for _, workSet := range exercise.WorkSets {
				workSet.Id = 0
				workSet.SetZeroTimes()
				workSet.ExerciseId = exercise.Id
				workSets = append(workSets, workSet)
			}
		}
	}

	err = w.WorkSetStore.InsertMany(&workSets)
	if err != nil {
		return err
	}

	return nil

}

func (w Week) duplicateWeekDays(templateWeekId, newWeekId int) (newWeekDays []model.WeekDay, templateWeekDays []model.WeekDay, err error) {
	err = w.WeekDayStore.DeleteByWeekId(newWeekId)
	if err != nil {
		// TODO: raise error and stop process
		return
	}

	newWeek, err := w.WeekStore.GetById(newWeekId)
	if err != nil {
		return
	}
	templateWeek, err := w.WeekStore.GetById(templateWeekId)
	if err != nil {
		return
	}
	dateDiffDays := newWeek.StartDate.Sub(templateWeek.StartDate).Hours() / 24

	templateWeekDays, err = w.WeekDayStore.GetByWeekIdWithDeleted(templateWeekId)
	if err != nil {
		return
	}

	newWeekDays = make([]model.WeekDay, len(templateWeekDays))
	for i, weekDay := range templateWeekDays {
		weekDay.Id = 0
		weekDay.SetZeroTimes()
		weekDay.DeletedAt = nil
		weekDay.WeekId = newWeekId
		weekDay.TimeslotId = nil
		weekDay.DayDate = weekDay.DayDate.AddDate(0, 0, int(dateDiffDays))
		newWeekDays[i] = weekDay
	}

	err = w.WeekDayStore.InsertMany(&newWeekDays)
	if err != nil {
		return
	}

	return
}
