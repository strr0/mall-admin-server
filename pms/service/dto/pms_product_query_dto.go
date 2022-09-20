package dto

type PmsProductQueryDto struct {
	PublishStatus     int32  `json:"publishStatus" form:"publishStatus"`
	VerifyStatus      int32  `json:"verifyStatus" form:"verifyStatus"`
	Keyword           string `json:"keyword" form:"keyword"`
	ProductSn         string `json:"productSn" form:"productSn"`
	ProductCategoryId int64  `json:"productCategoryId" form:"productCategoryId"`
	BrandId           int64  `json:"brandId" form:"brandId"`
}
