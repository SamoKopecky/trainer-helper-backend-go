package crud

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type TimeslotFull struct {
	model.Timeslot
	PersonName *string `json:"person_name"`
}

type CRUDTimeslot struct {
	Db *bun.DB
}

func (c CRUDTimeslot) GetByTimeRange(startDate, endDate time.Time) ([]TimeslotFull, error) {
	ctx := context.Background()
	var timeslots []TimeslotFull
	var err error

	c.Db.NewSelect().
		Model((*model.Timeslot)(nil)).
		ColumnExpr("person.name AS person_name").
		ColumnExpr("timeslot.*").
		Where("start BETWEEN ? AND ?", startDate, endDate).
		Join("LEFT JOIN person ON person.id = timeslot.user_id").
		Scan(ctx, &timeslots)

	return timeslots, err
}
