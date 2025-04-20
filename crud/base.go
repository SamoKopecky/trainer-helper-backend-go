package crud

import (
	"context"

	"github.com/uptrace/bun"
)

type CRUDBase[T any] struct {
	db bun.IDB
}

func (c CRUDBase[T]) Update(model *T) (err error) {
	query := c.db.NewUpdate().
		Model(model).
		OmitZero().
		WherePK()

	_, err = query.Exec(context.Background())
	return
}

func (c CRUDBase[T]) Insert(model *T) error {
	_, err := c.db.NewInsert().
		Model(model).
		Returning("*").
		Exec(context.Background())

	return err
}

func (c CRUDBase[T]) Get() (models []T, err error) {
	err = c.db.NewSelect().Model(&models).Scan(context.TODO())
	return
}

func (c CRUDBase[T]) InsertMany(models *[]T) error {
	if len(*models) == 0 {
		return nil
	}

	_, err := c.db.NewInsert().
		Model(models).
		Exec(context.Background())

	return err
}
