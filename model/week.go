package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Week struct {
	bun.BaseModel `bun:"table:week"`
	IdModel
	Timestamp
	DeletedTimestamp

	BlockId   int       `json:"block_id"`
	UserId    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	Label     int       `json:"label"`
	Note      *string   `json:"note"`

	// Not used in DB model
	WeekDays []WeekDay `bun:"rel:has-many,join:id=week_id" json:"-"`
}

func BuildWeek(blockId int, startDate time.Time, label int, userId string, note *string) *Week {
	return &Week{
		BlockId:   blockId,
		StartDate: startDate,
		Label:     label,
		UserId:    userId,
		Note:      note,
	}
}
