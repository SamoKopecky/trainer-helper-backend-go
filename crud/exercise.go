package crud

import (
	"context"
	"log"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDExercise struct {
	Db *bun.DB
}

func (c CRUDExercise) GetExerciseWorkSets(Id int32) ([]model.CRUDExerciseWorkSets, error) {
	ctx := context.Background()
	var res []model.CRUDExerciseWorkSets

	err := c.Db.NewSelect().
		Model((*model.Exercise)(nil)).
		ColumnExpr("exercise.timeslot_id, exercise.group_id, exercise.set_type, exercise.note, exercise.id AS exercise_id").
		ColumnExpr("work_set.exercise_id, work_set.reps, work_set.intensity, work_set.rpe, work_set.id AS work_set_id").
		Join("JOIN work_set ON work_set.exercise_id = exercise.id").
		Where("exercise.timeslot_id = ?", Id).
		Scan(ctx, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res, err
}
