package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Week struct {
	CRUDBase[model.Week]
}

func NewWeek(db bun.IDB) Week {
	return Week{CRUDBase: CRUDBase[model.Week]{db: db}}
}

func (w Week) GetByUserId(userId string) (weeks []model.Week, err error) {
	err = w.db.NewSelect().
		Model(&weeks).
		Where("user_id = ?", userId).
		Scan(context.Background())

	return
}
