package testutil

import (
	"trainer-helper/model"
	"trainer-helper/utils"
)

func BlockUserId(userId string) utils.FactoryOption[model.Block] {
	return func(b *model.Block) {
		b.UserId = userId
	}
}

func BlockId(id int) utils.FactoryOption[model.Block] {
	return func(b *model.Block) {
		b.Id = id
	}
}

func BlockLabel(label int) utils.FactoryOption[model.Block] {
	return func(b *model.Block) {
		b.Label = label
	}
}

func BlockFactory(options ...utils.FactoryOption[model.Block]) *model.Block {
	block := &model.Block{
		UserId: "1",
		Label:  utils.RandomInt(),
	}

	for _, option := range options {
		option(block)
	}
	return block
}
