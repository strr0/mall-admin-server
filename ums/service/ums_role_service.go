package service

import (
	"gorm.io/gen"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/util"
	"time"
)

type UmsRoleService struct {
	//
}

func (UmsRoleService) Create(umsRole model.UmsRole) error {
	umsRole.CreateTime = time.Now()
	return query.UmsRole.Create(&umsRole)
}

func (UmsRoleService) Update(idStr string, umsRole model.UmsRole) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsRole.Where(query.UmsRole.ID.Eq(id)).Updates(umsRole)
	return err
}

func (UmsRoleService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsRole.Where(query.UmsRole.ID.Eq(id)).Delete()
	return err
}

func (UmsRoleService) ListAll() []*model.UmsRole {
	find, err := query.UmsRole.Find()
	if err != nil {
		return nil
	}
	return find
}

func (UmsRoleService) List(keyword, pageStr, sizeStr string) ([]*model.UmsRole, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	umsRole := query.UmsRole
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds, umsRole.Name.Like(keyword))
	}
	umsRole.Where(conds...)
	total, err := umsRole.Count()
	if err != nil || total == 0 {
		return nil, 0
	}
	find, err := umsRole.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

// 修改角色状态
func (UmsRoleService) UpdateStatus(idStr, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	var umsRole model.UmsRole
	umsRole.Status = status
	_, err = query.UmsRole.Where(query.UmsRole.ID.Eq(id)).Updates(umsRole)
	return err
}

// 根据用户id获取角色
func (UmsRoleService) GetRoleList(adminIdStr string) []*model.UmsRole {
	adminId, err := util.ParseInt64WithErr(adminIdStr)
	if err != nil {
		return nil
	}
	list, err := query.UmsRole.GetRoleList(adminId)
	if err != nil {
		return nil
	}
	return list
}

// 根据用户id获取角色名称
func (UmsRoleService) GetRoleNameList(adminIdStr string) []string {
	adminId, err := util.ParseInt64WithErr(adminIdStr)
	if err != nil {
		return nil
	}
	list, err := query.UmsRole.GetRoleList(adminId)
	if err != nil {
		return nil
	}
	var res []string
	for _, item := range list {
		res = append(res, item.Name)
	}
	return res
}

// 根据角色id获取菜单
func (UmsRoleService) ListMenu(roleIdStr string) []*model.UmsMenu {
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return nil
	}
	menus, err := query.UmsMenu.GetMenuListByRoleId(roleId)
	if err != nil {
		return nil
	}
	return menus
}

// 根据角色id获取资源
func (UmsRoleService) ListResource(roleIdStr string) []*model.UmsResource {
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return nil
	}
	resources, err := query.UmsResource.GetResourceListByRoleId(roleId)
	if err != nil {
		return nil
	}
	return resources
}

// 菜单分配
func (UmsRoleService) AllocMenu(roleIdStr string, menuIdsStr []string) error {
	var err error
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return err
	}
	_, err = query.UmsRoleMenuRelation.Where(query.UmsRoleMenuRelation.RoleID.Eq(roleId)).Delete()
	if err != nil {
		return err
	}
	for _, menuIdStr := range menuIdsStr {
		menuId, err := util.ParseInt64WithErr(menuIdStr)
		if err == nil {
			rel := model.UmsRoleMenuRelation{RoleID: roleId, MenuID: menuId}
			err = query.UmsRoleMenuRelation.Create(&rel)
		}
	}
	return err
}

// 资源分配
func (UmsRoleService) AllocResource(roleIdStr string, resourceIdsStr []string) error {
	var err error
	roleId, err := util.ParseInt64WithErr(roleIdStr)
	if err != nil {
		return err
	}
	_, err = query.UmsRoleResourceRelation.Where(query.UmsRoleResourceRelation.RoleID.Eq(roleId)).Delete()
	if err != nil {
		return err
	}
	for _, resourceIdStr := range resourceIdsStr {
		resourceId, err := util.ParseInt64WithErr(resourceIdStr)
		if err == nil {
			rel := model.UmsRoleResourceRelation{RoleID: roleId, ResourceID: resourceId}
			err = query.UmsRoleResourceRelation.Create(&rel)
		}
	}
	return err
}
