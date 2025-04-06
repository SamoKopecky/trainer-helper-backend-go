package crud

import (
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type ExerciseType struct {
	CRUDBase[model.ExerciseType]
}

func NewExerciseType(db bun.IDB) ExerciseType {
	return ExerciseType{CRUDBase: CRUDBase[model.ExerciseType]{db: db}}
}
