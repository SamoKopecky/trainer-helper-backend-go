package crud

import (
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type WeekDay struct {
	CRUDBase[model.WeekDay]
}

func NewWeekDay(db bun.IDB) WeekDay {
	return WeekDay{CRUDBase: CRUDBase[model.WeekDay]{db: db}}
}
