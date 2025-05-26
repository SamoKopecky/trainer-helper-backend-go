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

func (et ExerciseType) UpdateMediaFile(id int, path, originalName string) error {
	_, err := et.db.NewUpdate().
		Model((*model.ExerciseType)(nil)).
		Set("file_path = ?", path).
		Set("original_file_name = ?", originalName).
		Where("id = ?", id).
		Exec(context.Background())

	return err

}
