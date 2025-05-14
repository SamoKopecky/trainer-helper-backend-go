package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Exercise struct {
	CRUDBase[model.Exercise]
}

func NewExercise(db bun.IDB) Exercise {
	return Exercise{CRUDBase: CRUDBase[model.Exercise]{db: db}}
}

func (e Exercise) GetExerciseWorkSets(weekDayIds []int) ([]model.Exercise, error) {
	ctx := context.Background()
	var res []model.Exercise

	err := e.db.NewSelect().
		Model(&res).
		Relation("WorkSets").
		Where("exercise.week_day_id IN (?)", bun.In(weekDayIds)).
		Scan(ctx)

	return res, err
}

func (e Exercise) DeleteByWeekDayId(weekDayId int) (err error) {
	_, err = e.db.NewDelete().
		Model((*model.Exercise)(nil)).
		Where("week_day_id = ?", weekDayId).
		Exec(context.Background())
	return
}
