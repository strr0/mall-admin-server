package service

import (
	"gorm.io/gen"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/util"
	"time"
)

type UmsAdminService struct {
	//
}

// 根据用户名获取用户
func (UmsAdminService) GetAdminByUsername(username string) *model.UmsAdmin {
	first, err := query.UmsAdmin.Where(query.UmsAdmin.Username.Eq(username)).First()
	if err != nil {
		return nil
	}
	return first
}

// 注册用户
func (UmsAdminService) Register(umsAdmin model.UmsAdmin) error {
	umsAdmin.CreateTime = time.Now()
	return query.UmsAdmin.Create(&umsAdmin)
}

// 登录
func (UmsAdminService) Login(username, password string) *model.UmsAdmin {
	first, err := query.UmsAdmin.Where(query.UmsAdmin.Username.Eq(username), query.UmsAdmin.Password.Eq(password)).First()
	if err != nil {
		return nil
	}
	return first
}

// 根据id获取用户
func (UmsAdminService) GetItem(idStr string) *model.UmsAdmin {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.UmsAdmin.Where(query.UmsAdmin.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

// 用户列表
func (UmsAdminService) List(keyword string, pageStr, sizeStr string) ([]*model.UmsAdmin, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	umsAdmin := query.UmsAdmin
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds, umsAdmin.NickName.Like(keyword))
	}
	umsAdmin.Where(conds...)
	total, err := umsAdmin.Count()
	if err != nil || total == 0 {
		return nil, 0
	}
	find, err := umsAdmin.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

// 更新用户信息
func (UmsAdminService) Update(idStr string, umsAdmin model.UmsAdmin) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsAdmin.Where(query.UmsAdmin.ID.Eq(id)).Updates(umsAdmin)
	return err
}

// 删除用户
func (UmsAdminService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsAdmin.Where(query.UmsAdmin.ID.Eq(id)).Delete()
	return err
}

// 更新状态
func (UmsAdminService) UpdateStatus(idStr, statusStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	status, err := util.ParseInt32WithErr(statusStr)
	if err != nil {
		return err
	}
	var umsAdmin model.UmsAdmin
	umsAdmin.Status = status
	_, err = query.UmsAdmin.Where(query.UmsAdmin.ID.Eq(id)).Updates(umsAdmin)
	return err
}

// 角色分配
func (UmsAdminService) UpdateRole(adminIdStr string, roleIdsStr []string) error {
	var err error
	adminId, err := util.ParseInt64WithErr(adminIdStr)
	if err != nil {
		return err
	}
	_, err = query.UmsAdminRoleRelation.Where(query.UmsAdminRoleRelation.AdminID.Eq(adminId)).Delete()
	if err != nil {
		return err
	}
	for _, roleIdStr := range roleIdsStr {
		roleId, err := util.ParseInt64WithErr(roleIdStr)
		if err == nil {
			rel := model.UmsAdminRoleRelation{AdminID: adminId, RoleID: roleId}
			err = query.UmsAdminRoleRelation.Create(&rel)
		}
	}
	return err
}
