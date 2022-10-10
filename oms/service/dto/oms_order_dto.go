package dto

import "mall-admin-server/oms/model"

type OmsOrderDto struct {
	model.OmsOrder
	OrderItemList []model.OmsOrderItem           `json:"orderItemList"`
	HistoryList   []model.OmsOrderOperateHistory `json:"historyList"`
}
