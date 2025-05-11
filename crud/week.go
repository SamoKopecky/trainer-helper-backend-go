package crud

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Week struct {
	CRUDBase[model.Week]
}

func NewWeek(db bun.IDB) Week {
	return Week{CRUDBase: CRUDBase[model.Week]{db: db}}
}

func (w Week) GetLastWeekDate(blockId int) (time.Time, error) {
	var week model.Week
	err := w.db.NewSelect().
		Model(&week).
		ColumnExpr("start_date").
		Where("block_id = ?", blockId).
		Order("label DESC").
		Limit(1).
		Scan(context.Background())

	if errors.Is(err, sql.ErrNoRows) {
		return week.StartDate, nil
	}

	return week.StartDate, err
}
