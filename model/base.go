package model

import "time"

type Timestamp struct {
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func buildTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type IdModel struct {
	Id int32 `bun:",pk,autoincrement" json:"id"`
}

func (im IdModel) IsEmpty() bool {
	return im.Id == 0
}
