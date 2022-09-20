package service

import (
	"fmt"
	"gorm.io/gen"
	cmsq "mall-admin-server/cms/query"
	pms "mall-admin-server/pms/model"
	pmsq "mall-admin-server/pms/query"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
	"time"
)

// 商品管理
type PmsProductService struct {
	//
}

func (iService PmsProductService) Create(pmsProductDto dto.PmsProductDto) error {
	pmsProduct := pmsProductDto.PmsProduct
	err := pmsq.PmsProduct.Create(&pmsProduct)
	if err != nil {
		return err
	}
	iService.insertRelations(pmsProduct.ID, pmsProductDto)
	return nil
}

func (PmsProductService) GetUpdateInfo(idStr string) *dto.PmsProductDto {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	pmsProductDto := &dto.PmsProductDto{
		PmsProduct: *first,
	}
	category, err := pmsq.PmsProductCategory.Where(pmsq.PmsProductCategory.ID.Eq(first.ProductCategoryID)).First()
	if err == nil {
		pmsProductDto.CateParentId = category.ParentID
	}
	find, err := pmsq.PmsProductLadder.Where(pmsq.PmsProductLadder.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.ProductLadderList = find
	}
	reductions, err := pmsq.PmsProductFullReduction.Where(pmsq.PmsProductFullReduction.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.ProductFullReductionList = reductions
	}
	prices, err := pmsq.PmsMemberPrice.Where(pmsq.PmsMemberPrice.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.MemberPriceList = prices
	}
	stocks, err := pmsq.PmsSkuStock.Where(pmsq.PmsSkuStock.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.SkuStockList = stocks
	}
	values, err := pmsq.PmsProductAttributeValue.Where(pmsq.PmsProductAttributeValue.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.ProductAttributeValueList = values
	}
	relations, err := cmsq.CmsSubjectProductRelation.Where(cmsq.CmsSubjectProductRelation.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.SubjectProductRelationList = relations
	}
	productRelations, err := cmsq.CmsPrefrenceAreaProductRelation.Where(cmsq.CmsPrefrenceAreaProductRelation.ProductID.Eq(id)).Find()
	if err == nil {
		pmsProductDto.PrefrenceAreaProductRelationList = productRelations
	}
	return pmsProductDto
}

func (iService PmsProductService) Update(idStr string, pmsProductDto dto.PmsProductDto) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	pmsProduct := pmsProductDto.PmsProduct
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.Eq(id)).Updates(pmsProduct)
	if err != nil {
		return err
	}
	// 阶梯价格
	_, err = pmsq.PmsProductLadder.Where(pmsq.PmsProductLadder.ProductID.Eq(id)).Delete()
	// 满减价格
	_, err = pmsq.PmsProductFullReduction.Where(pmsq.PmsProductFullReduction.ProductID.Eq(id)).Delete()
	// 会员价格
	_, err = pmsq.PmsMemberPrice.Where(pmsq.PmsMemberPrice.ProductID.Eq(id)).Delete()
	// 添加sku库存信息
	_, err = pmsq.PmsSkuStock.Where(pmsq.PmsSkuStock.ProductID.Eq(id)).Delete()
	// 添加自定义商品规格
	_, err = pmsq.PmsProductAttributeValue.Where(pmsq.PmsProductAttributeValue.ProductID.Eq(id)).Delete()
	// 关联专题
	_, err = cmsq.CmsSubjectProductRelation.Where(cmsq.CmsSubjectProductRelation.ProductID.Eq(id)).Delete()
	// 关联优选
	_, err = cmsq.CmsPrefrenceAreaProductRelation.Where(cmsq.CmsPrefrenceAreaProductRelation.ProductID.Eq(id)).Delete()
	iService.insertRelations(id, pmsProductDto)
	return nil
}

func (PmsProductService) List(pmsProductQueryDto dto.PmsProductQueryDto, pageStr, sizeStr string) ([]*pms.PmsProduct, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	pmsProduct := pmsq.PmsProduct
	conds := make([]gen.Condition, 0)
	conds = append(conds, pmsProduct.DeleteStatus.Eq(0))
	if pmsProductQueryDto.PublishStatus != 0 {
		conds = append(conds, pmsProduct.PublishStatus.Eq(pmsProductQueryDto.PublishStatus))
	}
	if pmsProductQueryDto.VerifyStatus != 0 {
		conds = append(conds, pmsProduct.VerifyStatus.Eq(pmsProductQueryDto.VerifyStatus))
	}
	if pmsProductQueryDto.Keyword != "" {
		conds = append(conds, pmsProduct.Name.Like("%"+pmsProductQueryDto.Keyword+"%"))
	}
	if pmsProductQueryDto.ProductSn != "" {
		conds = append(conds, pmsProduct.ProductSn.Eq(pmsProductQueryDto.ProductSn))
	}
	if pmsProductQueryDto.BrandId != 0 {
		conds = append(conds, pmsProduct.BrandID.Eq(pmsProductQueryDto.BrandId))
	}
	if pmsProductQueryDto.ProductCategoryId != 0 {
		conds = append(conds, pmsProduct.ProductCategoryID.Eq(pmsProductQueryDto.ProductCategoryId))
	}
	pmsProductDo := pmsProduct.Where(conds...)
	total, err := pmsProductDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := pmsProductDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

func (PmsProductService) UpdateVerifyStatus(idsStr []string, verifyStatusStr, detail string) error {
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
	var pmsProduct pms.PmsProduct
	pmsProduct.VerifyStatus = verifyStatus
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.In(ids...)).Updates(pmsProduct)
	if err != nil {
		return err
	}
	// 审核记录
	for _, id := range ids {
		var record pms.PmsProductVertifyRecord
		record.ProductID = id
		record.CreateTime = time.Now()
		record.Detail = detail
		record.Status = verifyStatus
		record.VertifyMan = "test"
		_ = pmsq.PmsProductVertifyRecord.Create(&record)
	}
	return nil
}

func (PmsProductService) UpdatePublishStatus(idsStr []string, publishStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	publishStatus, err := util.ParseInt32WithErr(publishStatusStr)
	if err != nil {
		return err
	}
	var pmsProduct pms.PmsProduct
	pmsProduct.PublishStatus = publishStatus
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.In(ids...)).Updates(pmsProduct)
	return err
}

func (PmsProductService) UpdateRecommendStatus(idsStr []string, recommendStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	recommendStatus, err := util.ParseInt32WithErr(recommendStatusStr)
	var pmsProduct pms.PmsProduct
	pmsProduct.RecommandStatus = recommendStatus
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.In(ids...)).Updates(pmsProduct)
	return err
}

func (PmsProductService) UpdateNewStatus(idsStr []string, newStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	newStatus, err := util.ParseInt32WithErr(newStatusStr)
	if err != nil {
		return err
	}
	var pmsProduct pms.PmsProduct
	pmsProduct.NewStatus = newStatus
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.In(ids...)).Updates(pmsProduct)
	return err
}

func (PmsProductService) UpdateDeleteStatus(idsStr []string, deleteStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	deleteStatus, err := util.ParseInt32WithErr(deleteStatusStr)
	var pmsProduct pms.PmsProduct
	pmsProduct.DeleteStatus = deleteStatus
	_, err = pmsq.PmsProduct.Where(pmsq.PmsProduct.ID.In(ids...)).Updates(pmsProduct)
	return err
}

func (PmsProductService) SimpleList(keyword string) []*pms.PmsProduct {
	pmsProduct := pmsq.PmsProduct
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds,
			pmsProduct.DeleteStatus.Eq(0),
			pmsProduct.Or(
				pmsProduct.Name.Like("%"+keyword+"%"),
				pmsProduct.ProductSn.Like("%"+keyword+"%"),
			),
		)
	}
	find, err := pmsProduct.Where(conds...).Find()
	if err != nil {
		return nil
	}
	return find
}

func (PmsProductService) insertRelations(productId int64, pmsProductDto dto.PmsProductDto) {
	// 阶梯价格
	for _, ladder := range pmsProductDto.ProductLadderList {
		ladder.ProductID = productId
		_ = pmsq.PmsProductLadder.Create(ladder)
	}
	// 满减价格
	for _, fullReduction := range pmsProductDto.ProductFullReductionList {
		fullReduction.ProductID = productId
		_ = pmsq.PmsProductFullReduction.Create(fullReduction)
	}
	// 会员价格
	for _, price := range pmsProductDto.MemberPriceList {
		price.ProductID = productId
		_ = pmsq.PmsMemberPrice.Create(price)
	}
	// 添加sku库存信息
	now := util.DateToStr(time.Now())
	for i, stock := range pmsProductDto.SkuStockList {
		stock.ProductID = productId
		if stock.SkuCode == "" {
			stock.SkuCode = fmt.Sprintf("%s%04d%03d", now, productId, i)
		}
		_ = pmsq.PmsSkuStock.Create(stock)
	}
	// 添加自定义商品规格
	for _, attributeValue := range pmsProductDto.ProductAttributeValueList {
		attributeValue.ProductID = productId
		_ = pmsq.PmsProductAttributeValue.Create(attributeValue)
	}
	// 关联专题
	for _, relation := range pmsProductDto.SubjectProductRelationList {
		relation.ProductID = productId
		_ = cmsq.CmsSubjectProductRelation.Create(relation)
	}
	// 关联优选
	for _, relation := range pmsProductDto.PrefrenceAreaProductRelationList {
		relation.ProductID = productId
		_ = cmsq.CmsPrefrenceAreaProductRelation.Create(relation)
	}
}
