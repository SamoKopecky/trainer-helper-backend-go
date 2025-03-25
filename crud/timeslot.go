package crud

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDTimeslot struct {
	CRUDBase[model.Timeslot]
}

func NewCRUDTimeslot(db *bun.DB) CRUDTimeslot {
	return CRUDTimeslot{CRUDBase: CRUDBase[model.Timeslot]{db: db}}
}

func (c CRUDTimeslot) GetByTimeRange(startDate, endDate time.Time) ([]*model.Timeslot, error) {
	ctx := context.Background()
	var timeslots []*model.Timeslot

	err := c.db.NewSelect().
		Model(&timeslots).
		Where("start BETWEEN ? AND ?", startDate, endDate).
		Scan(ctx)

	return timeslots, err
}

func (c CRUDTimeslot) GetById(timeslotId int32) (model.Timeslot, error) {
	ctx := context.Background()
	var timeslot model.Timeslot

	err := c.db.NewSelect().
		Model(&timeslot).
		Where("timeslot.id = ?", timeslotId).
		Scan(ctx)

	return timeslot, err

}

func (c CRUDTimeslot) Delete(timeslotId int32) error {
	ctx := context.Background()

	// Actually does soft delete
	_, err := c.db.NewDelete().
		Model((*model.Timeslot)(nil)).
		Where("id = ?", timeslotId).
		Exec(ctx)

	return err
}

func (c CRUDTimeslot) RevertSolfDelete(timeslotId int32) error {
	_, err := c.db.NewUpdate().
		Model((*model.Timeslot)(nil)).
		Set("deleted_at = ?", nil).
		WhereAllWithDeleted().
		Where("id = ?", timeslotId).
		Exec(context.Background())

	return err
}
