package schema

import "trainer-helper/model"

type BlocksMap map[int]BlocksValue

type BlocksValue struct {
	model.Block
	Weeks WeeksMap `json:"weeks"`
}
type WeeksMap map[int]model.Week
