package dto

type OmsUpdateStatusDto struct {
	Id               int64   `json:"id" form:"id"`
	CompanyAddressId int64   `json:"companyAddressId" form:"companyAddressId"`
	ReturnAmount     float64 `json:"returnAmount" form:"returnAmount"`
	HandleNote       string  `json:"handleNote" form:"handleNote"`
	HandleMan        string  `json:"handleMan" form:"handleMan"`
	ReceiveNote      string  `json:"receiveNote" form:"receiveNote"`
	ReceiveMan       string  `json:"receiveMan" form:"receiveMan"`
	Status           int32   `json:"status" form:"status"`
}
