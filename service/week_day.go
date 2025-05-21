package service

import (
	"trainer-helper/model"
	"trainer-helper/store"
)

type WeekDay struct {
	WeekDayStore  store.WeekDay
	ExerciseStore store.Exercise
	WorkSetStore  store.WorkSet
}
// TODO: Move this to week service
// on week day creation, also change date by using start date + 7 days
// don't forget to delete real all old entities for the new week

func (wd WeekDay) DuplicateWeekDays(weekDayIds []int, newWeekId int) error {
	weekDays, err := wd.duplicateWeekDays(weekDayIds, newWeekId)
	if err != nil {
		return err
	}
	newWeekDayIds := make([]int, len(weekDays))
	for i, weekDay := range weekDays {
		newWeekDayIds[i] = weekDay.Id
	}

	var workSets []model.WorkSet
	oldExercises, err := wd.ExerciseStore.GetExerciseWorkSets(newWeekDayIds)
	exercises := make([]model.Exercise, len(oldExercises))
	for i, exercise := range oldExercises {
		exercise.Id = 0
		exercise.SetZeroTimes()
		exercise.DeletedAt = nil
		exercise.WeekDayId =
	}

	return nil

}

func (wd WeekDay) duplicateWeekDays(weekDayIds []int, newWeekId int) ([]model.WeekDay, error) {
	oldWeekDays, err := wd.WeekDayStore.GetByWeekIdsWithDeleted(weekDayIds)
	if err != nil {
		return []model.WeekDay{}, err
	}
	newWeekDays := make([]model.WeekDay, len(oldWeekDays))
	for i, weekDay := range oldWeekDays {
		weekDay.Id = 0
		weekDay.SetZeroTimes()
		weekDay.DeletedAt = nil
		weekDay.WeekId = newWeekId
		newWeekDays[i] = weekDay
	}
	err = wd.WeekDayStore.InsertMany(&newWeekDays)
	if err != nil {
		return []model.WeekDay{}, err
	}
	err = wd.WeekDayStore.DeleteManyReal(weekDayIds)
	if err != nil {
		// Throw error and stop duplication
		return []model.WeekDay{}, err
	}
	return newWeekDays, nil
}
