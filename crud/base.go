package crud

import (
	"context"

	"github.com/uptrace/bun"
)

type CRUDBase[T any] struct {
	Db *bun.DB
}

func (c CRUDBase[T]) Update(model *T) error {
	ctx := context.Background()

	_, err := c.Db.NewUpdate().
		Model(model).
		OmitZero().
		WherePK().
		Exec(ctx)

	return err
}
