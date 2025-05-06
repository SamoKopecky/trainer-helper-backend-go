package block

import "trainer-helper/model"

type blockGetRequest struct {
	UserId string `query:"user_id"`
}

type blockPostRequest struct {
	Label int `json:"label"`
}

func (bpr blockPostRequest) toModel(userId string) *model.Block {
	return model.BuildBlock(userId, bpr.Label)
}

type blockDeleteRequest struct {
	Id int `json:"id"`
}
