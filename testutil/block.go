package testutil

import (
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func BlockUserId(t *testing.T, userId string) utils.FactoryOption[model.Block] {
	t.Helper()
	return func(b *model.Block) {
		b.UserId = userId
	}
}

func BlockId(t *testing.T, id int) utils.FactoryOption[model.Block] {
	t.Helper()
	return func(b *model.Block) {
		b.Id = id
	}
}

func BlockLabel(t *testing.T, label int) utils.FactoryOption[model.Block] {
	t.Helper()
	return func(b *model.Block) {
		b.Label = label
	}
}

func BlockFactory(t *testing.T, options ...utils.FactoryOption[model.Block]) *model.Block {
	t.Helper()
	block := &model.Block{
		UserId: "1",
		Label:  utils.RandomInt(),
	}
	block.Id = utils.RandomInt()

	for _, option := range options {
		option(block)
	}
	return block
}
