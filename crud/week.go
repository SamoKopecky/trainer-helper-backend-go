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

func (w Week) GetPreviousBlockId(userId string) (time.Time, error) {
	var week model.Week
	err := w.db.NewSelect().
		Model(&week).
		ColumnExpr("week.start_date").
		Join("JOIN block ON week.block_id = block.id and block.deleted_at IS NULL").
		Where("week.user_id = ?", userId).
		Order("block.label DESC", "week.label DESC").
		Limit(1).
		Scan(context.Background())

	if errors.Is(err, sql.ErrNoRows) {
		return week.StartDate, nil
	}

	return week.StartDate, err
}

func (w Week) GetClosestToDate(startDate time.Time, userId string) (model.Week, error) {
	var week model.Week
	err := w.db.NewRaw("SELECT * FROM week WHERE week.user_id = ? ORDER BY ? <-> week.start_date LIMIT 1", userId, startDate).
		Scan(context.Background(), &week)

	return week, err
}
