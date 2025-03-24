package crud

import (
	"context"

	"github.com/uptrace/bun"
)

type CRUDBase[T any] struct {
	db *bun.DB
}

func (c CRUDBase[T]) Update(modelT *T) error {
	_, err := c.db.NewUpdate().
		Model(modelT).
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
