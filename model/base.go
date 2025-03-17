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
