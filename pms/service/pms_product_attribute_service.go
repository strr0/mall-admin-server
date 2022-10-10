package service

import (
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品属性管理
type PmsProductAttributeService struct {
	DB *gorm.DB
}

func (iService PmsProductAttributeService) List(cid, type_, pageStr, sizeStr string) ([]model.PmsProductAttribute, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	where := iService.DB.Where("product_attribute_category_id = ? and type = ?", cid, type_)
	var count int64
	result := where.Model(&model.PmsProductAttribute{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.PmsProductAttribute
	result = where.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService PmsProductAttributeService) Create(pmsProductAttribute model.PmsProductAttribute) error {
	result := iService.DB.Create(&pmsProductAttribute)
	if result.Error != nil {
		return result.Error
	}
	// 更新分类
	var category model.PmsProductAttributeCategory
	result = iService.DB.First(&category, pmsProductAttribute.ProductAttributeCategoryID)
	if result.Error != nil {
		return result.Error
	}
	switch pmsProductAttribute.Type {
	case 0:
		result = iService.DB.Model(&category).Update("attribute_count", category.AttributeCount+1)
	case 1:
		result = iService.DB.Model(&category).Update("param_count", category.ParamCount+1)
	default:
	}
	return result.Error
}

func (iService PmsProductAttributeService) Update(idStr string, pmsProductAttribute model.PmsProductAttribute) error {
	result := iService.DB.Save(&pmsProductAttribute)
	return result.Error
}

func (iService PmsProductAttributeService) GetItem(id string) *model.PmsProductAttribute {
	var pmsProductAttribute model.PmsProductAttribute
	result := iService.DB.First(&pmsProductAttribute, id)
	if result.Error != nil {
		return nil
	}
	return &pmsProductAttribute
}

func (iService PmsProductAttributeService) Delete(ids []string) error {
	// 获取属性
	var pmsProductAttribute model.PmsProductAttribute
	result := iService.DB.First(&pmsProductAttribute, ids)
	if result.Error != nil {
		return result.Error
	}
	// 获取属性分类
	var pmsProductAttributeCategory model.PmsProductAttributeCategory
	result = iService.DB.First(&pmsProductAttributeCategory, pmsProductAttribute.ProductAttributeCategoryID)
	if result.Error != nil {
		return result.Error
	}
	// 删除属性
	result = iService.DB.Delete(&model.PmsProductAttribute{}, ids)
	if result.Error != nil {
		return result.Error
	}
	count := int32(result.RowsAffected)
	switch pmsProductAttribute.Type {
	case 0:
		var left int32
		if pmsProductAttributeCategory.AttributeCount > count {
			left = pmsProductAttributeCategory.AttributeCount - count
		} else {
			left = 0
		}
		result = iService.DB.Model(&pmsProductAttributeCategory).Update("attribute_count", left)
	case 1:
		var left int32
		if pmsProductAttributeCategory.ParamCount > count {
			left = pmsProductAttributeCategory.ParamCount - count
		} else {
			left = 0
		}
		result = iService.DB.Model(&pmsProductAttributeCategory).Update("param_count", left)
	default:
	}
	return result.Error
}

func (iService PmsProductAttributeService) GetProductAttrInfo(productCategoryId string) []dto.PmsProductAttributeDto {
	var list []dto.PmsProductAttributeDto
	result := iService.DB.Raw("SELECT pa.id attribute_id, pac.id attribute_category_id FROM pms_product_category_attribute_relation pcar LEFT JOIN pms_product_attribute pa ON pa.id = pcar.product_attribute_id LEFT JOIN pms_product_attribute_category pac ON pa.product_attribute_category_id = pac.id WHERE pcar.product_category_id = ?", productCategoryId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}
