package dto

import "time"

type OmsOrderQueryDto struct {
	OrderSn         string    `json:"orderSn" form:"orderSn"`
	ReceiverKeyword string    `json:"receiverKeyword" form:"receiverKeyword"`
	Status          int32     `json:"status" form:"status"`
	OrderType       int32     `json:"orderType" form:"orderType"`
	SourceType      int32     `json:"sourceType" form:"sourceType"`
	CreateTime      time.Time `json:"createTime" form:"createTime"`
}
