package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
	"time"
)

type UmsResourceCategoryService struct {
	DB *gorm.DB
}

// 获取所有分类
func (iService UmsResourceCategoryService) ListAll() []model.UmsResourceCategory {
	var list []model.UmsResourceCategory
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 创建分类
func (iService UmsResourceCategoryService) Create(umsResourceCategory model.UmsResourceCategory) error {
	umsResourceCategory.CreateTime = time.Now()
	result := iService.DB.Create(&umsResourceCategory)
	return result.Error
}

// 修改分类
func (iService UmsResourceCategoryService) Update(idStr string, umsResourceCategory model.UmsResourceCategory) error {
	result := iService.DB.Save(&umsResourceCategory)
	return result.Error
}

// 删除分类
func (iService UmsResourceCategoryService) Delete(id string) error {
	var umsResourceCategory model.UmsResourceCategory
	result := iService.DB.Delete(&umsResourceCategory, id)
	return result.Error
}
