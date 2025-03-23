package crud

import (
	"context"

	"github.com/uptrace/bun"
)

type CRUDBase[T any] struct {
	db *bun.DB
}

func (c CRUDBase[T]) Update(model *T) error {
	_, err := c.db.NewUpdate().
		Model(model).
		OmitZero().
		WherePK().
		Exec(context.Background())

	return err
}

func (c CRUDBase[T]) Insert(model *T) error {
	_, err := c.db.NewInsert().
		Model(model).
		Returning("*").
		Exec(context.Background())

	return err
}
