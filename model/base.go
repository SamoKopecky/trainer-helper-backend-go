package model

import "time"

type Timestamp struct {
	UpdatedAt time.Time
	CreatedAt time.Time
}

func buildTimestamp() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
