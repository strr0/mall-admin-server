package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
	"mall-admin-server/util"
	"time"
)

type UmsResourceService struct {
	DB *gorm.DB
}

// 创建资源
func (iService UmsResourceService) Create(umsResource model.UmsResource) error {
	umsResource.CreateTime = time.Now()
	result := iService.DB.Create(&umsResource)
	return result.Error
}

// 更新资源
func (iService UmsResourceService) Update(idStr string, umsResource model.UmsResource) error {
	result := iService.DB.Save(&umsResource)
	return result.Error
}

// 根据id获取资源
func (iService UmsResourceService) GetItem(id string) *model.UmsResource {
	var umsResource model.UmsResource
	result := iService.DB.First(&umsResource, id)
	if result.Error != nil {
		return nil
	}
	return &umsResource
}

// 删除资源
func (iService UmsResourceService) Delete(id string) error {
	var umsResource model.UmsResource
	result := iService.DB.Delete(&umsResource, id)
	return result.Error
}

// 资源列表
func (iService UmsResourceService) List(categoryId, nameKeyword, urlKeyword, pageStr, sizeStr string) ([]model.UmsResource, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if categoryId != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("category_id = ?", categoryId)
		})
	}
	if nameKeyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+nameKeyword+"%")
		})
	}
	if urlKeyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("url like ?", "%"+urlKeyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.UmsResource{}).Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []model.UmsResource
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

// 获取全部资源
func (iService UmsResourceService) ListAll() []model.UmsResource {
	var list []model.UmsResource
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 根据用户id获取资源
func (iService UmsResourceService) GetResourceList(adminId int64) []model.UmsResource {
	var list []model.UmsResource
	result := iService.DB.Raw("SELECT ur.id id, ur.create_time, ur.`name`, ur.url, ur.description, ur.category_id FROM ums_admin_role_relation ar LEFT JOIN ums_role_resource_relation rr ON ar.role_id = rr.role_id LEFT JOIN ums_resource ur ON ur.id = rr.resource_id WHERE ar.admin_id = ? AND ur.id IS NOT NULL GROUP BY ur.id", adminId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}
