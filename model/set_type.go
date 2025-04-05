package model

import (
	"github.com/uptrace/bun"
)

type SetType struct {
	bun.BaseModel `bun:"table:set_type"`
	IdModel

	UserId       string  `json:"user_id"`
	Name         string  `json:"name"`
	Note         *string `json:"note"`
	MediaType    *string `json:"media_type"`
	MediaAddress *string `json:"media_address"`
	Timestamp
}

func BuildSetType(userId, name string, note, mediaType, mediaAddress *string) *SetType {
	return &SetType{
		UserId:       userId,
		Name:         name,
		Note:         note,
		MediaType:    mediaType,
		MediaAddress: mediaAddress,
		Timestamp:    buildTimestamp(),
	}
}
