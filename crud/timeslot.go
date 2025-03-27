package crud

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Timeslot struct {
	CRUDBase[model.Timeslot]
}

func NewTimeslot(db *bun.DB) Timeslot {
	return Timeslot{CRUDBase: CRUDBase[model.Timeslot]{db: db}}
}

func (t Timeslot) GetByTimeRangeAndUserId(startDate, endDate time.Time, trainerId string, isTrainer bool) (timeslots []model.Timeslot, err error) {
	baseQuery := t.db.NewSelect().
		Model(&timeslots).
		Where("start BETWEEN ? AND ?", startDate, endDate)

	if isTrainer {
		baseQuery.Where("trainer_id = ?", trainerId)
	} else {
		baseQuery.Where("trainee_id = ?", trainerId)
	}

	err = baseQuery.Scan(context.Background())
	return
}

func (t Timeslot) GetById(timeslotId int32) (model.Timeslot, error) {
	ctx := context.Background()
	var timeslot model.Timeslot

	err := t.db.NewSelect().
		Model(&timeslot).
		Where("timeslot.id = ?", timeslotId).
		Scan(ctx)

	return timeslot, err

}

func (t Timeslot) Delete(timeslotId int32) error {
	ctx := context.Background()

	// Actually does soft delete
	_, err := t.db.NewDelete().
		Model((*model.Timeslot)(nil)).
		Where("id = ?", timeslotId).
		Exec(ctx)

	return err
}

func (t Timeslot) RevertSolfDelete(timeslotId int32) error {
	_, err := t.db.NewUpdate().
		Model((*model.Timeslot)(nil)).
		Set("deleted_at = ?", nil).
		WhereAllWithDeleted().
		Where("id = ?", timeslotId).
		Exec(context.Background())

	return err
}
