package store

import (
	"time"
	"trainer-helper/model"
)

type Week interface {
	StoreBase[model.Week]
	GetLastWeekDate(blockId int) (time.Time, error)
	GetPreviousBlockId(userId string) (time.Time, error)
	GetClosestToDate(startDate time.Time, userId string) (model.Week, error)
}
