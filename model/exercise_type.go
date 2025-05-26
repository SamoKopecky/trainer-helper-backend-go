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
	Timestamp

	UserId           string     `json:"user_id"`
	Name             string     `json:"name"`
	Note             *string    `json:"note"`
	MediaType        *MediaType `json:"media_type"`
	YoutubeLink      *string    `json:"youtube_link"`
	FilePath         *string    `json:"file_path"`
	OriginalFileName *string    `json:"original_file_name"`
}

func BuildExerciseType(userId, name string, note, youtubeLink, filePath, originalFileName *string, mediaType *MediaType) *ExerciseType {
	return &ExerciseType{
		UserId:           userId,
		Name:             name,
		Note:             note,
		MediaType:        mediaType,
		YoutubeLink:      youtubeLink,
		FilePath:         filePath,
		OriginalFileName: originalFileName,
		Timestamp:        buildTimestamp(),
	}
}
