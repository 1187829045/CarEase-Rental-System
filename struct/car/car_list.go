package car

import "car.rental/dao/model"

type CarListQuery struct {
	Status *int8 `form:"status"`
}
type CarListResp struct {
	Items  []*model.Car    `json:"items"`
	Counts map[int32]int64 `json:"counts"`
}
