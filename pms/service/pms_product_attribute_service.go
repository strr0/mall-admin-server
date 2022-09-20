package service

import (
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/query"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品属性管理
type PmsProductAttributeService struct {
	//
}

func (PmsProductAttributeService) List(cidStr, typeStr, pageStr, sizeStr string) ([]*model.PmsProductAttribute, int64) {
	cid, err := util.ParseInt64WithErr(cidStr)
	if err != nil {
		return nil, 0
	}
	type_, err := util.ParseInt32WithErr(typeStr)
	if err != nil {
		return nil, 0
	}
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	pmsProductAttributeDo := query.PmsProductAttribute.Where(
		query.PmsProductAttribute.ProductAttributeCategoryID.Eq(cid),
		query.PmsProductAttribute.Type.Eq(type_),
	)
	total, err := pmsProductAttributeDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := pmsProductAttributeDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

func (PmsProductAttributeService) Create(pmsProductAttribute model.PmsProductAttribute) error {
	err := query.PmsProductAttribute.Create(&pmsProductAttribute)
	if err != nil {
		return err
	}
	first, err := query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(pmsProductAttribute.ProductAttributeCategoryID)).First()
	if err != nil {
		return err
	}
	switch pmsProductAttribute.Type {
	case 0:
		first.AttributeCount = first.AttributeCount + 1
	case 1:
		first.ParamCount = first.ParamCount + 1
	default:
	}
	_, err = query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(pmsProductAttribute.ProductAttributeCategoryID)).Updates(first)
	return err
}

func (PmsProductAttributeService) Update(idStr string, pmsProductAttribute model.PmsProductAttribute) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.PmsProductAttribute.Where(query.PmsProductAttribute.ID.Eq(id)).Updates(pmsProductAttribute)
	return err
}

func (PmsProductAttributeService) GetItem(idStr string) *model.PmsProductAttribute {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.PmsProductAttribute.Where(query.PmsProductAttribute.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (PmsProductAttributeService) Delete(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	// 获取属性
	first, err := query.PmsProductAttribute.Where(query.PmsProductAttribute.ID.In(ids...)).First()
	if err != nil {
		return err
	}
	// 获取属性分类
	category, err := query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(first.ProductAttributeCategoryID)).First()
	if err != nil {
		return err
	}
	// 删除属性
	result, err := query.PmsProductAttribute.Where(query.PmsProductAttribute.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
	count := int32(result.RowsAffected)
	switch first.Type {
	case 0:
		if category.AttributeCount > count {
			category.AttributeCount = category.AttributeCount - count
		} else {
			category.AttributeCount = 0
		}
	case 1:
		if category.ParamCount > count {
			category.ParamCount = category.ParamCount - count
		} else {
			category.ParamCount = 0
		}
	default:
	}
	// 更新属性分类
	_, err = query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(first.ProductAttributeCategoryID)).Updates(category)
	return err
}

func (PmsProductAttributeService) GetProductAttrInfo(productCategoryIdStr string) []*dto.PmsProductAttributeDto {
	productCategoryId, err := util.ParseInt64WithErr(productCategoryIdStr)
	if err != nil {
		return nil
	}
	info, err := query.PmsProductAttribute.GetProductAttrInfo(productCategoryId)
	if err != nil {
		return nil
	}
	return info
}
