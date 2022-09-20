package dto

import (
	cms "mall-admin-server/cms/model"
	pms "mall-admin-server/pms/model"
)

type PmsProductDto struct {
	pms.PmsProduct
	CateParentId                     int64                                  `json:"cateParentId" form:"cateParentId"`
	ProductLadderList                []*pms.PmsProductLadder                `json:"productLadderList" form:"productLadderList"`
	ProductFullReductionList         []*pms.PmsProductFullReduction         `json:"productFullReductionList" form:"productFullReductionList"`
	MemberPriceList                  []*pms.PmsMemberPrice                  `json:"memberPriceList" form:"memberPriceList"`
	SkuStockList                     []*pms.PmsSkuStock                     `json:"skuStockList" form:"skuStockList"`
	ProductAttributeValueList        []*pms.PmsProductAttributeValue        `json:"productAttributeValueList" form:"productAttributeValueList"`
	SubjectProductRelationList       []*cms.CmsSubjectProductRelation       `json:"subjectProductRelationList" form:"subjectProductRelationList"`
	PrefrenceAreaProductRelationList []*cms.CmsPrefrenceAreaProductRelation `json:"prefrenceAreaProductRelationList" form:"prefrenceAreaProductRelationList"`
}
