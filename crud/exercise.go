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

func (c CRUDExercise) GetExerciseWorkSets(Id int32) ([]model.CRUDExerciseWorkSets, error) {
	ctx := context.Background()
	var res []model.CRUDExerciseWorkSets

	err := c.db.NewSelect().
		Model((*model.Exercise)(nil)).
		ColumnExpr("exercise.timeslot_id, exercise.group_id, exercise.set_type, exercise.note, exercise.id AS exercise_id").
		ColumnExpr("work_set.exercise_id, work_set.reps, work_set.intensity, work_set.rpe, work_set.id AS work_set_id").
		Join("JOIN work_set ON work_set.exercise_id = exercise.id").
		Where("exercise.timeslot_id = ?", Id).
		Scan(ctx, &res)

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
