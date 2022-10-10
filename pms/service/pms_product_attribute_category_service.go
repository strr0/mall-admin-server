package service

import (
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品属性分类管理
type PmsProductAttributeCategoryService struct {
	DB *gorm.DB
}

func (iService PmsProductAttributeCategoryService) Create(name string) error {
	var pmsProductAttributeCategory model.PmsProductAttributeCategory
	pmsProductAttributeCategory.Name = name
	result := iService.DB.Create(&pmsProductAttributeCategory)
	return result.Error
}

func (iService PmsProductAttributeCategoryService) Update(id, name string) error {
	result := iService.DB.Model(&model.PmsProductAttributeCategory{}).Where("id = ?", id).Update("name", name)
	return result.Error
}

func (iService PmsProductAttributeCategoryService) Delete(id string) error {
	result := iService.DB.Delete(&model.PmsProductAttributeCategory{}, id)
	return result.Error
}

func (iService PmsProductAttributeCategoryService) GetItem(id string) *model.PmsProductAttributeCategory {
	var pmsProductAttributeCategory model.PmsProductAttributeCategory
	result := iService.DB.First(&pmsProductAttributeCategory, id)
	if result.Error != nil {
		return nil
	}
	return &pmsProductAttributeCategory
}

func (iService PmsProductAttributeCategoryService) List(pageStr, sizeStr string) ([]model.PmsProductAttributeCategory, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	var count int64
	result := iService.DB.Model(&model.PmsProductAttributeCategory{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.PmsProductAttributeCategory
	result = iService.DB.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService PmsProductAttributeCategoryService) GetListWithAttr() []dto.PmsProductAttributeCategoryDto {
	var list []dto.PmsProductAttributeCategoryDto
	result := iService.DB.Raw("SELECT pac.id, pac.name, pa.id attr_id, pa.name attr_name FROM pms_product_attribute_category pac LEFT JOIN pms_product_attribute pa ON pac.id = pa.product_attribute_category_id AND pa.type=1").Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}
