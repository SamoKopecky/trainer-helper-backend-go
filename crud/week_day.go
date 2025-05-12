package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type WeekDay struct {
	CRUDBase[model.WeekDay]
}

func NewWeekDay(db bun.IDB) WeekDay {
	return WeekDay{CRUDBase: CRUDBase[model.WeekDay]{db: db}}
}

func (wd WeekDay) GetByWeekId(weekId int) (weekDays []model.WeekDay, err error) {
	err = wd.db.NewSelect().
		Model(&weekDays).
		Where("week_id = ?", weekId).
		Scan(context.Background())

	return
}
