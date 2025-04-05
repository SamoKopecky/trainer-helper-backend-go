package crud

import (
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type SetType struct {
	CRUDBase[model.SetType]
}

func NewSetType(db bun.IDB) SetType {
	return SetType{CRUDBase: CRUDBase[model.SetType]{db: db}}
}
