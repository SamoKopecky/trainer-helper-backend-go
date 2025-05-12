package crud

import (
	"context"
	"trainer-helper/model"

	"github.com/uptrace/bun"
)

type Block struct {
	CRUDBase[model.Block]
}

func NewBlock(db bun.IDB) Block {
	return Block{CRUDBase: CRUDBase[model.Block]{db: db}}
}

func (b Block) GetBlockWeeksByUserId(userId string) (blocks []model.Block, err error) {
	err = b.db.NewSelect().
		Model(&blocks).
		Relation("Weeks").
		Where("block.user_id = ?", userId).
		Scan(context.Background())

	return
}
