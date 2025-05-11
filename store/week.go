package store

import (
	"time"
	"trainer-helper/model"
)

type Week interface {
	StoreBase[model.Week]
	GetLastWeekDate(blockId int) (time.Time, error)
}
