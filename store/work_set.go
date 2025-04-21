package store

import "trainer-helper/model"

type WorkSet interface {
	StoreBase[model.WorkSet]
	DeleteMany(ids []int) (int, error)
	UpdateMany(models []model.WorkSet) error
}
