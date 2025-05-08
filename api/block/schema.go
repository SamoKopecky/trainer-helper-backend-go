package block

import "trainer-helper/model"

type blockGetRequest struct {
	UserId string `query:"user_id"`
}

type blockPostRequest struct {
	Label  int    `json:"label"`
	UserId string `json:"user_id"`
}

func (bpr blockPostRequest) toModel() *model.Block {
	return model.BuildBlock(bpr.UserId, bpr.Label)
}

type blockDeleteRequest struct {
	Id int `json:"id"`
}
