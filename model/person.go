package model

import (
	"github.com/uptrace/bun"
)

type Person struct {
	bun.BaseModel `bun:"table:person"`
	IdModel

	Name  string `json:"name"`
	Email string `json:"email"`
	Timestamp
}

func BuildPerson(name, email string) *Person {
	return &Person{
		Name:      name,
		Email:     email,
		Timestamp: buildTimestamp()}
}
