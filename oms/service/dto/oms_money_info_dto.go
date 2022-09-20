package dto

type OmsMoneyInfoDto struct {
	OrderId        int64   `json:"orderId" form:"orderId"`
	FreightAmount  float64 `json:"freightAmount" form:"freightAmount"`
	DiscountAmount float64 `json:"discountAmount" form:"discountAmount"`
	Status         int32   `json:"status" form:"status"`
}
