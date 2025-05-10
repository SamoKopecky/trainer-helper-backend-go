package model

import (
	"time"

	"github.com/uptrace/bun"
)

type WeekDay struct {
	bun.BaseModel `bun:"table:week_day"`
	IdModel
	Timestamp

	WeekId  int       `json:"week_id"`
	UserId  string    `json:"user_id"`
	DayDate time.Time `json:"day_date"`
	Name    *string   `json:"name"`
}

func BuildWeekDay(weekId int, userId string, DayDate time.Time, name *string) *WeekDay {
	return &WeekDay{
		DayDate: DayDate,
		Name:    name,
		UserId:  userId,
		WeekId:  weekId,
	}
}
