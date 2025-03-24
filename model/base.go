package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const EmptyId = 0

type Timestamp struct {
	bun.BaseModel

	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func buildTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Timeslot) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		t.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		t.UpdatedAt = time.Now()
	}
	return nil
}

type IdModel struct {
	bun.BaseModel

	Id int32 `bun:",pk,autoincrement" json:"id"`
}

func (im IdModel) IsEmpty() bool {
	return im.Id == EmptyId
}
