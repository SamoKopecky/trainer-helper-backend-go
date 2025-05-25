package service

import (
	"testing"
	"time"
	"trainer-helper/model"
	store "trainer-helper/store/mock"
	"trainer-helper/testutil"
	"trainer-helper/utils"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWeek_NotFirst(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := time.Date(2025, 05, 12, 0, 0, 0, 0, time.UTC)
	service := Week{WeekStore: w}
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(nextMonday, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, nextMonday, 1, "1", nil)
	// Act
	if err := service.CreateWeek(newWeek, false); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// +7 to get to the next monday
	assert.Equal(t, nextMonday.AddDate(0, 0, 7), newWeek.StartDate)
}

func TestCreateWeek__NoLastWeek__NotFirst(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := utils.GetNextMonday(time.Now())
	service := Week{WeekStore: w}
	// Rerturn 0 time
	w.EXPECT().GetLastWeekDate(mock.Anything).Return(time.Time{}, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, time.Time{}, 1, "1", nil)
	// Act
	if err := service.CreateWeek(newWeek, false); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// 7 to get to the next monday
	assert.Equal(t, nextMonday.Day(), newWeek.StartDate.Day())
}

func TestCreateWeek__First(t *testing.T) {
	// Arrange
	w := store.NewMockWeek(t)
	nextMonday := time.Date(2025, 05, 12, 0, 0, 0, 0, time.UTC)
	service := Week{WeekStore: w}
	w.EXPECT().GetPreviousBlockId(mock.Anything).Return(nextMonday, nil)
	w.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Week) error {
		model1.IdModel = model.IdModel{
			Id: 10,
		}
		return nil
	})

	newWeek := model.BuildWeek(1, nextMonday, 1, "1", nil)
	// Act
	if err := service.CreateWeek(newWeek, true); err != nil {
		t.Fatalf("Failed to create weeks: %v", err)
	}

	// Assert
	w.Mock.AssertNumberOfCalls(t, "Insert", 1)
	// +7 to get to the next monday
	assert.Equal(t, nextMonday.AddDate(0, 0, 7), newWeek.StartDate)
}

func TestDuplicateWeekDays(t *testing.T) {
	w := store.NewMockWeek(t)
	wd := store.NewMockWeekDay(t)
	service := Week{WeekStore: w, WeekDayStore: wd}
	now := time.Now()
	week := testutil.WeekFactory(t, testutil.WeekDate(t, now))

	templateWeek := testutil.WeekFactory(t)
	weekDays := make([]model.WeekDay, 7)
	for i := range 7 {
		weekDays[i] = *testutil.WeekDayFactory(t,
			testutil.WeekDayIds(t, "1", templateWeek.Id))
	}
	timelostId := 2
	weekDays[0].TimeslotId = &timelostId

	wd.EXPECT().DeleteByWeekId(week.Id).Return(nil)
	w.EXPECT().GetById(week.Id).Return(*week, nil)
	wd.EXPECT().GetByWeekIdWithDeleted(templateWeek.Id).Return(weekDays, nil)

	var insertedArgs []model.WeekDay
	wd.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WeekDay) error {
		insertedArgs = *models
		return nil
	})
	createWeekDays, templateWeekDays, err := service.duplicateWeekDays(templateWeek.Id, week.Id)

	assert.Nil(t, err)
	assert.Equal(t, weekDays, templateWeekDays)

	for i := range 7 {
		weekDays[i].Id = 0
		weekDays[i].WeekId = week.Id
		weekDays[i].DayDate = now.AddDate(0, 0, i)
		weekDays[i].TimeslotId = nil
	}

	assert.Equal(t, weekDays, insertedArgs)
	assert.Equal(t, weekDays, createWeekDays)
}

func TestDuplicateWeekDaysPublic(t *testing.T) {
	w := store.NewMockWeek(t)
	wd := store.NewMockWeekDay(t)
	e := store.NewMockExercise(t)
	ws := store.NewMockWorkSet(t)
	service := Week{WeekStore: w, WeekDayStore: wd, ExerciseStore: e, WorkSetStore: ws}
	now := time.Now()

	week := testutil.WeekFactory(t, testutil.WeekDate(t, now))
	templateWeek := testutil.WeekFactory(t)

	weekDays := make([]model.WeekDay, 7)
	weekDayIds := make([]int, 7)
	for i := range 7 {
		weekDays[i] = *testutil.WeekDayFactory(t,
			testutil.WeekDayIds(t, "1", templateWeek.Id))
		weekDayIds[i] = weekDays[i].Id
	}

	exercises := make([]model.Exercise, 2)
	exercises[0] = *testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, weekDayIds[0]))
	exercises[1] = *testutil.ExerciseFactory(t, testutil.ExerciseWeekDayId(t, weekDayIds[1]))

	workSets := make([]model.WorkSet, 2)

	workSets[0] = *testutil.WorkSetFactory(t, testutil.WorkSetExerciseId(t, exercises[0].Id))
	exercises[0].WorkSets = append(exercises[0].WorkSets, workSets[0])

	workSets[1] = *testutil.WorkSetFactory(t, testutil.WorkSetExerciseId(t, exercises[1].Id))
	exercises[1].WorkSets = append(exercises[1].WorkSets, workSets[1])

	var expectedExerciseInserts []model.Exercise
	var expectedWorkSetInsert []model.WorkSet
	var newWeekDayIds []int

	wd.EXPECT().DeleteByWeekId(week.Id).Return(nil)
	w.EXPECT().GetById(week.Id).Return(*week, nil)
	wd.EXPECT().GetByWeekIdWithDeleted(templateWeek.Id).Return(weekDays, nil)
	wd.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WeekDay) error {
		for i, _ := range *models {
			(*models)[i].Id = utils.RandomInt()
			newWeekDayIds = append(newWeekDayIds, (*models)[i].Id)
		}
		return nil
	})
	e.EXPECT().GetExerciseWorkSets(weekDayIds).Return(exercises, nil)
	e.EXPECT().Insert(mock.Anything).RunAndReturn(func(model1 *model.Exercise) error {
		// Simluate insert
		model1.Id = utils.RandomInt()
		expectedExerciseInserts = append(expectedExerciseInserts, *model1)
		return nil
	}).Times(2)
	ws.EXPECT().InsertMany(mock.Anything).RunAndReturn(func(models *[]model.WorkSet) error {
		expectedWorkSetInsert = *models
		return nil
	})

	service.DuplicateWeekDays(templateWeek.Id, week.Id)

	ignoreFields := cmpopts.IgnoreFields(model.Exercise{}, "Id", "Timestamp", "WeekDayId")
	assert.True(t, cmp.Equal(expectedExerciseInserts, exercises, ignoreFields),
		cmp.Diff(expectedExerciseInserts, exercises, ignoreFields))
	assert.Equal(t, expectedExerciseInserts[0].WeekDayId, newWeekDayIds[0])
	assert.Equal(t, expectedExerciseInserts[1].WeekDayId, newWeekDayIds[1])

	ignoreFields = cmpopts.IgnoreFields(model.WorkSet{}, "Id", "Timestamp", "ExerciseId")
	assert.True(t, cmp.Equal(expectedWorkSetInsert, workSets, ignoreFields), cmp.Diff(expectedWorkSetInsert, workSets, ignoreFields))
	assert.Equal(t, expectedWorkSetInsert[0].ExerciseId, expectedExerciseInserts[0].Id)
	assert.Equal(t, expectedWorkSetInsert[1].ExerciseId, expectedExerciseInserts[1].Id)

}
