package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDExercise struct {
	CRUDBase[model.Exercise]
}

func NewCRUDExercise(db *bun.DB) CRUDExercise {
	return CRUDExercise{CRUDBase: CRUDBase[model.Exercise]{db: db}}
}

func (c CRUDExercise) GetExerciseWorkSets(Id int32) ([]*model.Exercise, error) {
	ctx := context.Background()
	var res []*model.Exercise

	err := c.db.NewSelect().
		Model(&res).
		Relation("WorkSets").
		Where("exercise.timeslot_id = ?", Id).
		Scan(ctx)

	return res, err
}

func (c CRUDExercise) DeleteByExerciseAndTimeslot(timeslotId, exerciseId int32) error {
	_, err := c.db.NewDelete().
		Model((*model.Exercise)(nil)).
		Where("timeslot_id = ?", timeslotId).
		Where("id = ?", exerciseId).
		Exec(context.Background())
	return err
}

func (c CRUDExercise) DeleteByTimeslot(timeslotId int32) (err error) {
	_, err = c.db.NewDelete().
		Model((*model.Exercise)(nil)).
		Where("timeslot_id = ?", timeslotId).
		Exec(context.Background())
	return
}
