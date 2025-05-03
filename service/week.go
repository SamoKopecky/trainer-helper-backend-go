package service

import (
	"trainer-helper/schema"
	"trainer-helper/store"
)

type Week struct {
	Store store.Week
}

func (w Week) GetBlocks(userId string) (blocks schema.Blocks, err error) {
	blocks = make(schema.Blocks)
	// weeks, err := w.Store.GetByUserId(userId)
	if err != nil {
		return
	}

	// for _, week := range weeks {
	// 	blocks[week.] = append(blocks[week.BlockLabel], week)
	// }

	return
}
