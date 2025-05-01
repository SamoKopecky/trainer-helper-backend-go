package store

import "trainer-helper/model"

type Week interface {
	StoreBase[model.Week]
	GetByUserId(userId string) (weeks []model.Week, err error)
}
