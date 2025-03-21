package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type CRUDPerson struct {
	Db *bun.DB
}

func (c CRUDPerson) Get(id int32) (model.Person, error) {
	ctx := context.Background()
	var person model.Person

	err := c.Db.NewSelect().Model(&person).Where("id = ?", id).Scan(ctx)
	return person, err
}
