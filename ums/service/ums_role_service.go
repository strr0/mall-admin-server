package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
	"mall-admin-server/util"
	"time"
)

type UmsRoleService struct {
	DB *gorm.DB
}

func (iService UmsRoleService) Create(umsRole model.UmsRole) error {
	umsRole.CreateTime = time.Now()
	result := iService.DB.Create(&umsRole)
	return result.Error
}

func (iService UmsRoleService) Update(idStr string, umsRole model.UmsRole) error {
	result := iService.DB.Save(&umsRole)
	return result.Error
}

func (iService UmsRoleService) Delete(id string) error {
	result := iService.DB.Delete(&model.UmsRole{}, id)
	return result.Error
}

func (iService UmsRoleService) ListAll() []model.UmsRole {
	var list []model.UmsRole
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService UmsRoleService) List(keyword, pageStr, sizeStr string) ([]model.UmsRole, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+keyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.UmsRole{}).Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []model.UmsRole
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

// 修改角色状态
func (iService UmsRoleService) UpdateStatus(id, status string) error {
	result := iService.DB.Model(&model.UmsRole{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

// 根据用户id获取角色
func (iService UmsRoleService) GetRoleList(adminId string) []model.UmsRole {
	var list []model.UmsRole
	result := iService.DB.Raw("select r.* from ums_admin_role_relation ar left join ums_role r on ar.role_id = r.id where ar.admin_id = ?", adminId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 根据用户id获取角色名称
func (iService UmsRoleService) GetRoleNameList(adminId string) []string {
	var list []model.UmsRole
	result := iService.DB.Raw("select r.* from ums_admin_role_relation ar left join ums_role r on ar.role_id = r.id where ar.admin_id = ?", adminId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	var res []string
	for _, item := range list {
		res = append(res, item.Name)
	}
	return res
}

// 根据角色id获取菜单
func (iService UmsRoleService) ListMenu(roleId string) []model.UmsMenu {
	var list []model.UmsMenu
	result := iService.DB.Raw("SELECT m.id, m.parent_id, m.create_time, m.title, m.level, m.sort, m.name, m.icon, m.hidden FROM ums_role_menu_relation rmr LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE rmr.role_id = ? AND m.id IS NOT NULL GROUP BY m.i", roleId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 根据角色id获取资源
func (iService UmsRoleService) ListResource(roleId string) []model.UmsResource {
	var list []model.UmsResource
	result := iService.DB.Raw("SELECT r.id, r.create_time, r.`name`, r.url, r.description, r.category_id FROM ums_role_resource_relation rrr LEFT JOIN ums_resource r ON rrr.resource_id = r.id WHERE rrr.role_id = ? AND r.id IS NOT NULL GROUP BY r.id", roleId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 菜单分配
func (iService UmsRoleService) AllocMenu(roleIdStr string, menuIdsStr []string) error {
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return err
	}
	result := iService.DB.Where("role_id = ?", roleId).Delete(&model.UmsRoleMenuRelation{})
	if result.Error != nil {
		return err
	}
	for _, menuIdStr := range menuIdsStr {
		menuId, err := util.ParseInt64WithErr(menuIdStr)
		if err == nil {
			rel := model.UmsRoleMenuRelation{RoleID: roleId, MenuID: menuId}
			result = iService.DB.Create(&rel)
		}
	}
	return result.Error
}

// 资源分配
func (iService UmsRoleService) AllocResource(roleIdStr string, resourceIdsStr []string) error {
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return err
	}
	result := iService.DB.Where("role_id = ?", roleId).Delete(&model.UmsRoleResourceRelation{})
	if result.Error != nil {
		return err
	}
	for _, resourceIdStr := range resourceIdsStr {
		resourceId, err := util.ParseInt64WithErr(resourceIdStr)
		if err == nil {
			rel := model.UmsRoleResourceRelation{RoleID: roleId, ResourceID: resourceId}
			result = iService.DB.Create(&rel)
		}
	}
	return result.Error
}
