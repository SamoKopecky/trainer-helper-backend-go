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

func (ws WorkSet) UpdateMany(models []model.WorkSet) error {
	values := ws.db.NewValues(&models)

	_, err := ws.db.NewUpdate().
		With("_data", values).
		Model((*model.WorkSet)(nil)).
		TableExpr("_data").
		Set("rpe = _data.rpe").
		Set("reps = _data.reps").
		Set("intensity = _data.intensity").
		Where("work_set.id = _data.id").
		Exec(context.Background())
	return err
}
