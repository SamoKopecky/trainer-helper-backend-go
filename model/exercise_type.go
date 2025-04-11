package model

import (
	"github.com/uptrace/bun"
)

type MediaType string

const (
	Youtube MediaType = "YOUTUBE"
	File    MediaType = "FILE"
)

type ExerciseType struct {
	bun.BaseModel `bun:"table:exercise_type"`
	IdModel

	UserId       string     `json:"user_id"`
	Name         string     `json:"name"`
	Note         *string    `json:"note"`
	MediaType    *MediaType `json:"media_type"`
	MediaAddress *string    `json:"media_address"`
	Timestamp
}

func BuildExerciseType(userId, name string, note, mediaAddress *string, mediaType *MediaType) *ExerciseType {
	return &ExerciseType{
		UserId:       userId,
		Name:         name,
		Note:         note,
		MediaType:    mediaType,
		MediaAddress: mediaAddress,
		Timestamp:    buildTimestamp(),
	}
}
