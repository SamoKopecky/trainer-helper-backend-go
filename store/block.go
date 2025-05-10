package store

import "trainer-helper/model"

type Block interface {
	StoreBase[model.Block]
	GetBlockWeeksByUserId(userId string) (blocks []model.Block, err error)
}
