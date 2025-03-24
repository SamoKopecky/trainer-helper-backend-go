package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDWorkSet struct {
	CRUDBase[model.WorkSet]
}

func NewCRUDWorkSet(db *bun.DB) CRUDWorkSet {
	return CRUDWorkSet{CRUDBase: CRUDBase[model.WorkSet]{db: db}}
}

func (c CRUDWorkSet) InsertMany(work_sets *[]model.WorkSet) error {
	_, err := c.db.NewInsert().
		Model(work_sets).
		Exec(context.Background())

	return err
}

func (c CRUDWorkSet) DeleteMany(ids []int32) (int32, error) {
	var deletedModels []model.WorkSet
	_, err := c.db.NewDelete().
		Model(&deletedModels).
		Returning("id").
		Where("id IN (?)", bun.In(ids)).
		Exec(context.Background())

	return int32(len(deletedModels)), err
}
