package crud

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDTimeslot struct {
	Db *bun.DB
}

func (c CRUDTimeslot) GetByTimeRange(startDate, endDate time.Time) ([]model.TimeslotFull, error) {
	ctx := context.Background()
	var timeslots []model.TimeslotFull
	var err error

	err = c.Db.NewSelect().
		Model((*model.Timeslot)(nil)).
		ColumnExpr("person.name AS person_name").
		ColumnExpr("timeslot.*").
		Where("start BETWEEN ? AND ?", startDate, endDate).
		Join("LEFT JOIN person ON person.id = timeslot.user_id").
		Scan(ctx, &timeslots)

	return timeslots, err
}

func (c CRUDTimeslot) Insert(timeslot *model.Timeslot) error {
	ctx := context.Background()

	_, err := c.Db.NewInsert().
		Model(timeslot).
		Exec(ctx)

	return err
}

func (c CRUDTimeslot) Delete(timeslotId int32) (*model.Timeslot, error) {
	ctx := context.Background()

	var timeslot model.Timeslot
	_, err := c.Db.NewDelete().
		Model(&timeslot).
		Where("id = ?", timeslotId).
		Returning("*").
		Exec(ctx)

	return &timeslot, err
}

func (c CRUDTimeslot) Update(timeslot *model.Timeslot) error {
	ctx := context.Background()

	_, err := c.Db.NewUpdate().
		Model(timeslot).
		OmitZero().
		WherePK().
		Exec(ctx)

	return err
}
