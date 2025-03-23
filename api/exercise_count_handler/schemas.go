package exercise_count_handler

import "trainer-helper/model"

type exerciseCountPostParams struct {
	Count           int32         `json:"count"`
	WorkSetTemplate model.WorkSet `json:"work_set_template"`
}

type exerciseCountDeleteParams struct {
	WorkSetIds []int32 `json:"work_set_ids"`
}
