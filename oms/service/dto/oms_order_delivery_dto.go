package dto

type OmsOrderDeliveryDto struct {
	OrderId         int64  `json:"orderId" form:"orderId"`
	DeliveryCompany string `json:"deliveryCompany" form:"deliveryCompany"`
	DeliverySn      string `json:"deliverySn" form:"deliverySn"`
}
