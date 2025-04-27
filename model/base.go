package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const EmptyId = 0

type Timestamp struct {
	bun.BaseModel

	UpdatedAt time.Time  `json:"-"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" bun:",soft_delete,nullzero"`
}

func (t *Timestamp) SetZeroTimes() {
	t.UpdatedAt = time.Time{}
	t.CreatedAt = time.Time{}
}

func buildTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
}

func (t *Timestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		t.CreatedAt = time.Now().UTC()
	case *bun.UpdateQuery:
		t.UpdatedAt = time.Now().UTC()
	}
	return nil
}

type IdModel struct {
	bun.BaseModel

	Id int `bun:",pk,autoincrement" json:"id"`
}

func (im IdModel) IsEmpty() bool {
	return im.Id == EmptyId
}
