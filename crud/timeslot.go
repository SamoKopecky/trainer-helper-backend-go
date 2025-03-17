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

func (c CRUDTimeslot) GetByTimeRange(startDate, endDate time.Time) ([]model.Timeslot, error) {
	ctx := context.Background()
	var timeslots []model.Timeslot
	var err error

	c.Db.NewSelect().
		Model(&timeslots).
		Where("start between ? and ?", startDate, endDate).
		Scan(ctx)

	return timeslots, err
}
