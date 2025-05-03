package model

import (
	"time"

	"github.com/uptrace/bun"
)

type WeekDay struct {
	bun.BaseModel `bun:"table:week_day"`
	IdModel
	Timestamp

	WeekId  int        `json:"week_id"`
	UserId  string     `json:"user_id"`
	Name    string     `json:"name"`
	DayDate *time.Time `json:"day_date"`
}

func BuildWeekDay(weekId int, userId, name string, DayDate *time.Time) *WeekDay {
	return &WeekDay{
		DayDate: DayDate,
		Name:    name,
		UserId:  userId,
		WeekId:  weekId,
	}
}
