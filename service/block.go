package service

import (
	"cmp"
	"slices"
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

	slices.SortFunc(blocks, func(a, b model.Block) int {
		return cmp.Compare(a.Label, b.Label)
	})
	for i := range blocks {
		if len(blocks[i].Weeks) == 0 {
			blocks[i].Weeks = []model.Week{}
		} else {
			slices.SortFunc(blocks[i].Weeks, func(a, b model.Week) int {
				return cmp.Compare(a.Label, b.Label)
			})
			for j := range blocks[i].Weeks {
				if len(blocks[i].Weeks[j].WeekDays) == 0 {
					blocks[i].Weeks[j].WeekDays = []model.WeekDay{}
				} else {
					slices.SortFunc(blocks[i].Weeks[j].WeekDays, func(a, b model.WeekDay) int {
						if a.DayDate.Before(b.DayDate) {
							return -1
						}
						if a.DayDate.After(b.DayDate) {
							return 1
						}
						return 0
					})
				}
			}
		}
	}

	return
}
