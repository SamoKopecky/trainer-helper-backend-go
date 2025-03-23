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

func (c CRUDTimeslot) getTimeslotQuery() *bun.SelectQuery {
	return c.db.NewSelect().
		Model((*model.Timeslot)(nil)).
		ColumnExpr("person.name AS person_name").
		ColumnExpr("timeslot.*").
		Join("LEFT JOIN person ON person.id = timeslot.user_id")
}

func (c CRUDTimeslot) GetByTimeRange(startDate, endDate time.Time) ([]model.ApiTimeslot, error) {
	ctx := context.Background()
	var timeslots []model.ApiTimeslot

	err := c.getTimeslotQuery().
		Where("start BETWEEN ? AND ?", startDate, endDate).
		Scan(ctx, &timeslots)

	return timeslots, err
}

func (c CRUDTimeslot) GetById(timeslotId int32) (model.ApiTimeslot, error) {
	ctx := context.Background()
	var timeslot model.ApiTimeslot

	err := c.getTimeslotQuery().
		Where("timeslot.id = ?", timeslotId).
		Scan(ctx, &timeslot)

	return timeslot, err

}

func (c CRUDTimeslot) Delete(timeslotId int32) (*model.Timeslot, error) {
	ctx := context.Background()

	var timeslot model.Timeslot
	_, err := c.db.NewDelete().
		Model(&timeslot).
		Where("id = ?", timeslotId).
		Returning("*").
		Exec(ctx)

	return &timeslot, err
}
