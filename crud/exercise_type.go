package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type ExerciseType struct {
	CRUDBase[model.ExerciseType]
}

func NewExerciseType(db bun.IDB) ExerciseType {
	return ExerciseType{CRUDBase: CRUDBase[model.ExerciseType]{db: db}}
}

func (et ExerciseType) GetByUserId(userId string) (res []model.ExerciseType, err error) {
	err = et.db.NewSelect().
		Model(&res).
		Where("user_id = ?", userId).
		Scan(context.Background())

	return
}

func (et ExerciseType) Undelete(modelId int) error {
	return ErrNotImplemented
}
