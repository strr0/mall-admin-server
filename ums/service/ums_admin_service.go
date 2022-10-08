package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
	"mall-admin-server/util"
	"time"
)

type UmsAdminService struct {
	DB *gorm.DB
}

// 根据用户名获取用户
func (iService UmsAdminService) GetAdminByUsername(username string) *model.UmsAdmin {
	var umsAdmin model.UmsAdmin
	result := iService.DB.Where("username = ?", username).First(&umsAdmin)
	if result.Error != nil {
		return nil
	}
	return &umsAdmin
}

// 注册用户
func (iService UmsAdminService) Register(umsAdmin model.UmsAdmin) error {
	umsAdmin.CreateTime = time.Now()
	result := iService.DB.Create(&umsAdmin)
	return result.Error
}

// 登录
func (iService UmsAdminService) Login(username, password string) *model.UmsAdmin {
	var umsAdmin model.UmsAdmin
	result := iService.DB.Where("username = ? and password = ?", username, password).First(&umsAdmin)
	if result.Error != nil {
		return nil
	}
	return &umsAdmin
}

// 根据id获取用户
func (iService UmsAdminService) GetItem(id string) *model.UmsAdmin {
	var umsAdmin model.UmsAdmin
	result := iService.DB.First(&umsAdmin, id)
	if result.Error != nil {
		return nil
	}
	return &umsAdmin
}

// 用户列表
func (iService UmsAdminService) List(keyword string, pageStr, sizeStr string) ([]model.UmsAdmin, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("nickname like ?", "%"+keyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.UmsAdmin{}).Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []model.UmsAdmin
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

// 更新用户信息
func (iService UmsAdminService) Update(idStr string, umsAdmin model.UmsAdmin) error {
	result := iService.DB.Save(&umsAdmin)
	return result.Error
}

// 删除用户
func (iService UmsAdminService) Delete(id string) error {
	result := iService.DB.Delete(&model.UmsAdmin{}, id)
	return result.Error
}

// 更新状态
func (iService UmsAdminService) UpdateStatus(id, status string) error {
	result := iService.DB.Model(&model.UmsAdmin{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

// 角色分配
func (iService UmsAdminService) UpdateRole(adminIdStr string, roleIdsStr []string) error {
	var err error
	adminId, err := util.ParseInt64WithErr(adminIdStr)
	if err != nil {
		return err
	}
	result := iService.DB.Delete(&model.UmsAdmin{}, adminId)
	if result.Error != nil {
		return result.Error
	}
	var relation model.UmsAdminRoleRelation
	result = iService.DB.Where("admin_id", adminId).Delete(&relation)
	if result.Error != nil {
		return result.Error
	}
	for _, roleIdStr := range roleIdsStr {
		roleId, err := util.ParseInt64WithErr(roleIdStr)
		if err == nil {
			rel := model.UmsAdminRoleRelation{AdminID: adminId, RoleID: roleId}
			result = iService.DB.Create(&rel)
		}
	}
	return result.Error
}
