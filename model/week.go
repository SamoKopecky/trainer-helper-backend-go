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
