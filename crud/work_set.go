package crud

import (
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDWorkSet struct {
	CRUDBase[model.WorkSet]
}

func NewCRUDWorkSet(db *bun.DB) CRUDWorkSet {
	return CRUDWorkSet{CRUDBase: CRUDBase[model.WorkSet]{Db: db}}
}
