package dto

type OmsReceiverInfoDto struct {
	OrderId               int64  `json:"orderId" form:"orderId"`
	ReceiverName          string `json:"receiverName" form:"receiverName"`
	ReceiverPhone         string `json:"receiverPhone" form:"receiverPhone"`
	ReceiverPostCode      string `json:"receiverPostCode" form:"receiverPostCode"`
	ReceiverDetailAddress string `json:"receiverDetailAddress" form:"receiverDetailAddress"`
	ReceiverProvince      string `json:"receiverProvince" form:"receiverProvince"`
	ReceiverCity          string `json:"receiverCity" form:"receiverCity"`
	ReceiverRegion        string `json:"receiverRegion" form:"receiverRegion"`
	Status                int32  `json:"status" form:"status"`
}
