package crud

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

var ErrNotImplemented = errors.New("This store is not implemented for this model")

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

func (c CRUDBase[T]) GetById(modelId int) (model T, err error) {
	err = c.db.NewSelect().
		Model(&model).
		Where("id = ?", modelId).
		Scan(context.TODO())
	return
}

func (c CRUDBase[T]) Delete(modelId int) error {
	ctx := context.Background()

	// Actually does soft delete
	_, err := c.db.NewDelete().
		Model((*T)(nil)).
		Where("id = ?", modelId).
		Exec(ctx)

	return err
}

func (c CRUDBase[T]) DeleteMany(modelIds []int) error {
	ctx := context.Background()

	// Actually does soft delete
	_, err := c.db.NewDelete().
		Model((*T)(nil)).
		Where("id IN (?)", bun.In(modelIds)).
		Exec(ctx)

	return err
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

func (c CRUDBase[T]) UndeleteMany(modelIds []int) error {
	_, err := c.db.NewUpdate().
		Model((*T)(nil)).
		Set("deleted_at = ?", nil).
		WhereAllWithDeleted().
		Where("id IN (?)", bun.In(modelIds)).
		Exec(context.Background())

	return err
}
