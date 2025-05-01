package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Week struct {
	bun.BaseModel `bun:"table:week"`
	IdModel
	Timestamp

	UserId     string    `json:"user_id"`
	StartDate  time.Time `json:"start"`
	Label      int       `json:"label"`
	BlockLabel int       `json:"block_label"`
	Monday     *string   `json:"monday"`
	Tuesday    *string   `json:"tuesday"`
	Wednesday  *string   `json:"wednesday"`
	Thursday   *string   `json:"thursday"`
	Friday     *string   `json:"friday"`
	Saturday   *string   `json:"saturday"`
	Sunday     *string   `json:"sunday"`
}

func BuildWeek(userId string, startDate time.Time, label, blockLabel int, monday, tuesday, wednesday, thursday, friday, saturday, sunday *string) *Week {
	return &Week{
		UserId:     userId,
		StartDate:  startDate,
		Label:      label,
		BlockLabel: blockLabel,
		Monday:     monday,
		Tuesday:    tuesday,
		Wednesday:  wednesday,
		Thursday:   thursday,
		Friday:     friday,
		Saturday:   saturday,
		Sunday:     sunday,
	}
}
