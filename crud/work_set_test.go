package crud

import (
	"context"
	"testing"
	"trainer-helper/model"
	"trainer-helper/utils"
)

func TestInsertMany(t *testing.T) {
	db := testSetup(t)
	crud := NewWorkSet(db)
	model3 := model.WorkSet{}
	crud.Insert(&model3)
	utils.PrettyPrint(model3)

	var model2 []model.WorkSet
	err := db.NewSelect().Model(&model2).Scan(context.TODO())
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(model2)
}
