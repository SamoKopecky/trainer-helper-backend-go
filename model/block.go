package model

import (
	"github.com/uptrace/bun"
)

type Block struct {
	bun.BaseModel `bun:"table:block"`
	IdModel
	Timestamp
	DeletedTimestamp

	UserId string `json:"user_id"`
	Label  int    `json:"label"`

	// Not used in DB model
	Weeks []Week `bun:"rel:has-many,join:id=block_id" json:"weeks"`
}

func BuildBlock(userId string, label int) *Block {
	return &Block{
		UserId: userId,
		Label:  label,
	}
}
