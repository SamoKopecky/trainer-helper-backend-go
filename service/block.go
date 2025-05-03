package service

import (
	"trainer-helper/schema"
	"trainer-helper/store"
)

type Block struct {
	Store store.Block
}

func (b Block) GetBlocks(userId string) (blocksMap schema.BlocksMap, err error) {
	blocksMap = make(schema.BlocksMap)
	blocks, err := b.Store.GetBlockWeeksByUserId(userId)
	if err != nil {
		return
	}

	for _, block := range blocks {
		weeksMap := make(schema.WeeksMap)
		for _, week := range block.Weeks {
			weeksMap[week.Label] = week
		}
		blocksMap[block.Label] = schema.BlocksValue{
			Block: block,
			Weeks: weeksMap,
		}
	}

	return
}
