package service

import (
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品分类管理
type PmsProductCategoryService struct {
	DB *gorm.DB
}

func (iService PmsProductCategoryService) Create(pmsProductCategoryDto dto.PmsProductCategoryDto) error {
	pmsProductCategory := pmsProductCategoryDto.PmsProductCategory
	pmsProductCategory.ProductCount = 0
	iService.setCategoryLevel(&pmsProductCategory)
	result := iService.DB.Create(&pmsProductCategory)
	if result.Error != nil {
		return result.Error
	}
	// 创建关联属性
	if len(pmsProductCategoryDto.ProductAttributeIdList) > 0 {
		err := iService.insertRelationList(pmsProductCategoryDto.ID, pmsProductCategoryDto.ProductAttributeIdList)
		if err != nil {
			return err
		}
	}
	return nil
}

func (iService PmsProductCategoryService) Update(id string, pmsProductCategoryDto dto.PmsProductCategoryDto) error {
	pmsProductCategory := pmsProductCategoryDto.PmsProductCategory
	iService.setCategoryLevel(&pmsProductCategory)
	// 更新商品的分类名称
	result := iService.DB.Model(&model.PmsProduct{}).Where("product_category_id = ?", pmsProductCategoryDto.ID).Update("product_category_name", pmsProductCategory.Name)
	if result.Error != nil {
		return result.Error
	}
	// 更新关联属性
	result = iService.DB.Where("product_category_id = ?", pmsProductCategoryDto.ID).Delete(&model.PmsProductCategoryAttributeRelation{})
	if result.Error != nil {
		return result.Error
	}
	if len(pmsProductCategoryDto.ProductAttributeIdList) > 0 {
		err := iService.insertRelationList(pmsProductCategoryDto.ID, pmsProductCategoryDto.ProductAttributeIdList)
		if err != nil {
			return err
		}
	}
	result = iService.DB.Save(&pmsProductCategory)
	return result.Error
}

func (iService PmsProductCategoryService) List(parentId, pageStr, sizeStr string) ([]model.PmsProductCategory, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	where := iService.DB.Where("parent_id = ?", parentId)
	var count int64
	result := where.Model(&model.PmsProductCategory{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.PmsProductCategory
	result = where.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService PmsProductCategoryService) Delete(id string) error {
	result := iService.DB.Delete(&model.PmsProductCategory{}, id)
	return result.Error
}

func (iService PmsProductCategoryService) GetItem(id string) *model.PmsProductCategory {
	var pmsProductCategory model.PmsProductCategory
	result := iService.DB.First(&pmsProductCategory, id)
	if result.Error != nil {
		return nil
	}
	return &pmsProductCategory
}

func (iService PmsProductCategoryService) UpdateNavStatus(ids []string, navStatus string) error {
	result := iService.DB.Model(&model.PmsProductCategory{}).Where("id in ?", ids).Update("nav_status", navStatus)
	return result.Error
}

func (iService PmsProductCategoryService) UpdateShowStatus(ids []string, showStatus string) error {
	result := iService.DB.Model(&model.PmsProductCategory{}).Where("id in ?", ids).Update("show_status", showStatus)
	return result.Error
}

func (iService PmsProductCategoryService) TreeList() []*dto.PmsProductCategoryTreeDto {
	var list []model.PmsProductCategory
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	categoryMap := make(map[int64]*dto.PmsProductCategoryTreeDto)
	res := make([]*dto.PmsProductCategoryTreeDto, 0)
	for _, category := range list {
		categoryDto := dto.PmsProductCategoryTreeDto{
			PmsProductCategory: category,
			Children:           make([]dto.PmsProductCategoryTreeDto, 0),
		}
		categoryMap[category.ID] = &categoryDto
		if category.ParentID == 0 {
			res = append(res, &categoryDto)
		} else {
			if parent, exist := categoryMap[category.ParentID]; exist {
				parent.Children = append(parent.Children, categoryDto)
			}
		}
	}
	return res
}

// 创建关联属性
func (iService PmsProductCategoryService) insertRelationList(productCategoryId int64, productAttributeIdList []int64) error {
	for _, productAttributeId := range productAttributeIdList {
		rel := model.PmsProductCategoryAttributeRelation{
			ProductAttributeID: productAttributeId,
			ProductCategoryID:  productCategoryId,
		}
		_ = iService.DB.Create(&rel)
	}
	return nil
}

// 设置分类level
func (iService PmsProductCategoryService) setCategoryLevel(pmsProductCategory *model.PmsProductCategory) {
	if pmsProductCategory.ParentID == 0 {
		pmsProductCategory.Level = 0
	} else {
		var parent model.PmsProductCategory
		result := iService.DB.First(&parent, pmsProductCategory.ParentID)
		if result.Error != nil {
			return
		}
		if parent.ID != 0 {
			pmsProductCategory.Level = parent.Level + 1
		} else {
			pmsProductCategory.Level = 0
		}
	}
}
