package store

import "trainer-helper/model"

type WorkSet interface {
	StoreBase[model.WorkSet]
	UpdateMany(models []model.WorkSet) error
}
