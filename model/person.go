package model

import (
	"github.com/uptrace/bun"
)

type Person struct {
	bun.BaseModel `bun:"table:person"`
	IdModel

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
