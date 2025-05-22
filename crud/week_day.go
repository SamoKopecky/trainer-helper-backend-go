package crud

import (
	"context"
	"time"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type WeekDay struct {
	CRUDBase[model.WeekDay]
}

func NewWeekDay(db bun.IDB) WeekDay {
	return WeekDay{CRUDBase: CRUDBase[model.WeekDay]{db: db}}
}

func (wd WeekDay) GetByWeekIdWithDeleted(weekId int) (weekDays []model.WeekDay, err error) {
	err = wd.db.NewSelect().
		Model(&weekDays).
		WhereAllWithDeleted().
		Where("week_id = ?", weekId).
		Order("day_date DESC").
		Scan(context.Background())

	return
}

func (wd WeekDay) GetByTimeslotIds(timeslotIds []int) (weekDays []model.WeekDay, err error) {
	err = wd.db.NewSelect().
		Model(&weekDays).
		Where("timeslot_id IN (?)", bun.In(timeslotIds)).
		Scan(context.Background())

	return
}

func (wd WeekDay) GetByDate(dayDate time.Time, userId string) (weekDays []model.WeekDay, err error) {
	dateString := dayDate.Format("2006-01-02")

	err = wd.db.NewSelect().
		Model(&weekDays).
		Where("day_date = ? AND user_id = ?", dateString, userId).
		Scan(context.Background())

	return

}

func (wd WeekDay) DeleteTimeslot(weekId int) error {
	_, err := wd.db.NewUpdate().
		Model((*model.WeekDay)(nil)).
		Set("timeslot_id = NULL").
		Where("id = ?", weekId).
		Exec(context.Background())

	return err
}
func (wd WeekDay) DeleteByWeekId(weekId int) error {
	err := wd.db.NewDelete().
		Model((*model.WeekDay)(nil)).
		Where("week_day = ?", weekId).
		Scan(context.Background())

	return err
}
