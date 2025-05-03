package service

import (
	"trainer-helper/model"
	"trainer-helper/store"
)

type Block struct {
	Store store.Block
}

func (b Block) GetBlocks(userId string) (blocks []model.Block, err error) {
	blocks, err = b.Store.GetBlockWeeksByUserId(userId)
	if err != nil {
		return
	}

	for i := range blocks {
		if len(blocks[i].Weeks) == 0 {
			blocks[i].Weeks = []model.Week{}
		} else {
			for j := range blocks[i].Weeks {
				if len(blocks[i].Weeks[j].WeekDays) == 0 {
					blocks[i].Weeks[j].WeekDays = []model.WeekDay{}
				}
			}
		}
	}

	return
}
