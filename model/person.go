package model

import (
	"github.com/uptrace/bun"
)

type Person struct {
	bun.BaseModel `bun:"table:person"`

	Id    int32 `bun:",pk,autoincrement"`
	Name  string
	Email string
	Timestamp
}

func BuildPerson(name, email string) *Person {
	return &Person{
		Name:      name,
		Email:     email,
		Timestamp: buildTimestamp()}
}
