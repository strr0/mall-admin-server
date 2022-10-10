package service

import (
	"fmt"
	"gorm.io/gorm"
	cms "mall-admin-server/cms/model"
	pms "mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"time"
)

// 商品管理
type PmsProductService struct {
	DB *gorm.DB
}

func (iService PmsProductService) Create(pmsProductDto dto.PmsProductDto) error {
	pmsProduct := pmsProductDto.PmsProduct
	result := iService.DB.Create(&pmsProduct)
	if result.Error != nil {
		return result.Error
	}
	iService.insertRelations(pmsProduct.ID, pmsProductDto)
	return nil
}

func (iService PmsProductService) GetUpdateInfo(id string) *dto.PmsProductDto {
	var pmsProduct pms.PmsProduct
	result := iService.DB.First(&pmsProduct, id)
	if result.Error != nil {
		return nil
	}
	pmsProductDto := &dto.PmsProductDto{
		PmsProduct: pmsProduct,
	}
	var category pms.PmsProductCategory
	result = iService.DB.First(&category, pmsProduct.ProductCategoryID)
	if result.Error == nil {
		pmsProductDto.CateParentId = category.ParentID
	}
	var ladderList []pms.PmsProductLadder
	result = iService.DB.Where("product_id = ?", id).Find(&ladderList)
	if result.Error == nil {
		pmsProductDto.ProductLadderList = ladderList
	}
	var reductionList []pms.PmsProductFullReduction
	result = iService.DB.Where("product_id = ?", id).Find(&reductionList)
	if result.Error == nil {
		pmsProductDto.ProductFullReductionList = reductionList
	}
	var priceList []pms.PmsMemberPrice
	result = iService.DB.Where("product_id = ?", id).Find(&priceList)
	if result.Error == nil {
		pmsProductDto.MemberPriceList = priceList
	}
	var stockList []pms.PmsSkuStock
	result = iService.DB.Where("product_id = ?", id).Find(&stockList)
	if result.Error == nil {
		pmsProductDto.SkuStockList = stockList
	}
	var valueList []pms.PmsProductAttributeValue
	result = iService.DB.Where("product_id = ?", id).Find(&valueList)
	if result.Error == nil {
		pmsProductDto.ProductAttributeValueList = valueList
	}
	var relationList []cms.CmsSubjectProductRelation
	result = iService.DB.Where("product_id = ?", id).Find(&relationList)
	if result.Error == nil {
		pmsProductDto.SubjectProductRelationList = relationList
	}
	var areaList []cms.CmsPrefrenceAreaProductRelation
	result = iService.DB.Where("product_id = ?", id).Find(&areaList)
	if result.Error == nil {
		pmsProductDto.PrefrenceAreaProductRelationList = areaList
	}
	return pmsProductDto
}

func (iService PmsProductService) Update(id string, pmsProductDto dto.PmsProductDto) error {
	result := iService.DB.Save(&pmsProductDto.PmsProduct)
	if result.Error != nil {
		return result.Error
	}
	// 阶梯价格
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&pms.PmsProductLadder{})
	// 满减价格
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&pms.PmsProductFullReduction{})
	// 会员价格
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&pms.PmsMemberPrice{})
	// 添加sku库存信息
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&pms.PmsSkuStock{})
	// 添加自定义商品规格
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&pms.PmsProductAttributeValue{})
	// 关联专题
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&cms.CmsSubjectProductRelation{})
	// 关联优选
	result = iService.DB.Where("product_id = ?", pmsProductDto.ID).Delete(&cms.CmsPrefrenceAreaProductRelation{})
	iService.insertRelations(pmsProductDto.ID, pmsProductDto)
	return nil
}

func (iService PmsProductService) List(pmsProductQueryDto dto.PmsProductQueryDto, pageStr, sizeStr string) ([]pms.PmsProduct, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
		return db.Where("delete_status = ?", 0)
	})
	if pmsProductQueryDto.PublishStatus != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("publish_status = ?", pmsProductQueryDto.PublishStatus)
		})
	}
	if pmsProductQueryDto.VerifyStatus != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("verify_status = ?", pmsProductQueryDto.VerifyStatus)
		})
	}
	if pmsProductQueryDto.Keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+pmsProductQueryDto.Keyword+"%")
		})
	}
	if pmsProductQueryDto.ProductSn != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("product_sn = ?", pmsProductQueryDto.ProductSn)
		})
	}
	if pmsProductQueryDto.BrandId != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("brand_id = ?", pmsProductQueryDto.BrandId)
		})
	}
	if pmsProductQueryDto.ProductCategoryId != 0 {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("product_category_id = ?", pmsProductQueryDto.ProductCategoryId)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&pms.PmsProduct{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []pms.PmsProduct
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService PmsProductService) UpdateVerifyStatus(idsStr []string, verifyStatusStr, detail string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	verifyStatus, err := util.ParseInt32WithErr(verifyStatusStr)
	if err != nil {
		return err
	}
	result := iService.DB.Model(&pms.PmsProduct{}).Where("id in ?", ids).Update("verify_status", verifyStatus)
	if result.Error != nil {
		return result.Error
	}
	// 审核记录
	var records []pms.PmsProductVertifyRecord
	for _, id := range ids {
		var record pms.PmsProductVertifyRecord
		record.ProductID = id
		record.CreateTime = time.Now()
		record.Detail = detail
		record.Status = verifyStatus
		record.VertifyMan = "test"
		records = append(records, record)
	}
	result = iService.DB.Create(&records)
	return result.Error
}

func (iService PmsProductService) UpdatePublishStatus(ids []string, publishStatus string) error {
	result := iService.DB.Model(&pms.PmsProduct{}).Where("id in ?", ids).Update("publish_status", publishStatus)
	return result.Error
}

func (iService PmsProductService) UpdateRecommendStatus(ids []string, recommendStatus string) error {
	result := iService.DB.Model(&pms.PmsProduct{}).Where("id in ?", ids).Update("recommend_status", recommendStatus)
	return result.Error
}

func (iService PmsProductService) UpdateNewStatus(ids []string, newStatus string) error {
	result := iService.DB.Model(&pms.PmsProduct{}).Where("id in ?", ids).Update("new_status", newStatus)
	return result.Error
}

func (iService PmsProductService) UpdateDeleteStatus(ids []string, deleteStatus string) error {
	result := iService.DB.Model(&pms.PmsProduct{}).Where("id in ?", ids).Update("delete_status", deleteStatus)
	return result.Error
}

func (iService PmsProductService) SimpleList(keyword string) []pms.PmsProduct {
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("delete_status = ? and (name like ? or product_sn like ?)", 0, "%"+keyword+"%", "%"+keyword+"%")
		})
	}
	var list []pms.PmsProduct
	result := iService.DB.Scopes(funcs...).Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService PmsProductService) insertRelations(productId int64, pmsProductDto dto.PmsProductDto) {
	// 阶梯价格
	ladderList := make([]pms.PmsProductLadder, 0)
	for _, ladder := range pmsProductDto.ProductLadderList {
		ladder.ProductID = productId
		ladderList = append(ladderList, ladder)
	}
	iService.DB.Create(&ladderList)
	// 满减价格
	fullReductionList := make([]pms.PmsProductFullReduction, 0)
	for _, fullReduction := range pmsProductDto.ProductFullReductionList {
		fullReduction.ProductID = productId
		fullReductionList = append(fullReductionList, fullReduction)
	}
	iService.DB.Create(&fullReductionList)
	// 会员价格
	priceList := make([]pms.PmsMemberPrice, 0)
	for _, price := range pmsProductDto.MemberPriceList {
		price.ProductID = productId
		priceList = append(priceList, price)
	}
	iService.DB.Create(&priceList)
	// 添加sku库存信息
	stockList := make([]pms.PmsSkuStock, 0)
	now := util.DateToStr(time.Now())
	for i, stock := range pmsProductDto.SkuStockList {
		stock.ProductID = productId
		if stock.SkuCode == "" {
			stock.SkuCode = fmt.Sprintf("%s%04d%03d", now, productId, i)
		}
		stockList = append(stockList, stock)
	}
	iService.DB.Create(&stockList)
	// 添加自定义商品规格
	attributeValueList := make([]pms.PmsProductAttributeValue, 0)
	for _, attributeValue := range pmsProductDto.ProductAttributeValueList {
		attributeValue.ProductID = productId
		attributeValueList = append(attributeValueList, attributeValue)
	}
	iService.DB.Create(&attributeValueList)
	// 关联专题
	relationList := make([]cms.CmsSubjectProductRelation, 0)
	for _, relation := range pmsProductDto.SubjectProductRelationList {
		relation.ProductID = productId
		relationList = append(relationList, relation)
	}
	iService.DB.Create(&relationList)
	// 关联优选
	areaList := make([]cms.CmsPrefrenceAreaProductRelation, 0)
	for _, relation := range pmsProductDto.PrefrenceAreaProductRelationList {
		relation.ProductID = productId
		areaList = append(areaList, relation)
	}
	iService.DB.Create(&areaList)
}
