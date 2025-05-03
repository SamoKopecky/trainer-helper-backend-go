package crud

import (
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Week struct {
	CRUDBase[model.Week]
}

func NewWeek(db bun.IDB) Week {
	return Week{CRUDBase: CRUDBase[model.Week]{db: db}}
}
