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
	_, err := c.db.NewInsert().Model(work_sets).Exec(context.Background())

	return err

}
