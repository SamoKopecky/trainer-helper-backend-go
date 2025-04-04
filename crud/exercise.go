package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Exercise struct {
	CRUDBase[model.Exercise]
}

func NewExercise(db *bun.DB) Exercise {
	return Exercise{CRUDBase: CRUDBase[model.Exercise]{db: db}}
}

func (e Exercise) GetExerciseWorkSets(Id int) ([]*model.Exercise, error) {
	ctx := context.Background()
	var res []*model.Exercise

	err := e.db.NewSelect().
		Model(&res).
		Relation("WorkSets").
		Where("exercise.timeslot_id = ?", Id).
		Scan(ctx)

	return res, err
}

func (e Exercise) DeleteByExerciseAndTimeslot(timeslotId, exerciseId int) error {
	_, err := e.db.NewDelete().
		Model((*model.Exercise)(nil)).
		Where("timeslot_id = ?", timeslotId).
		Where("id = ?", exerciseId).
		Exec(context.Background())
	return err
}

func (e Exercise) DeleteByTimeslot(timeslotId int) (err error) {
	_, err = e.db.NewDelete().
		Model((*model.Exercise)(nil)).
		Where("timeslot_id = ?", timeslotId).
		Exec(context.Background())
	return
}
