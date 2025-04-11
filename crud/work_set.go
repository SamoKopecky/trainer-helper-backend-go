package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type WorkSet struct {
	CRUDBase[model.WorkSet]
}

func NewWorkSet(db bun.IDB) WorkSet {
	return WorkSet{CRUDBase: CRUDBase[model.WorkSet]{db: db}}
}

func (ws WorkSet) DeleteMany(ids []int) (int, error) {
	var deletedModels []model.WorkSet
	_, err := ws.db.NewDelete().
		Model(&deletedModels).
		Returning("id").
		Where("id IN (?)", bun.In(ids)).
		Exec(context.Background())

	return int(len(deletedModels)), err
}
