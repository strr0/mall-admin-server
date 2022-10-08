package service

import (
	"gorm.io/gorm"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/service/dto"
	"mall-admin-server/util"
	"time"
)

type UmsMenuService struct {
	DB *gorm.DB
}

// 创建菜单
func (iService UmsMenuService) Create(umsMenu model.UmsMenu) error {
	umsMenu.CreateTime = time.Now()
	result := iService.DB.Create(&umsMenu)
	return result.Error
}

// 更新菜单
func (iService UmsMenuService) Update(idStr string, umsMenu model.UmsMenu) error {
	result := iService.DB.Save(&umsMenu)
	return result.Error
}

// 根据id获取菜单
func (iService UmsMenuService) GetItem(id string) *model.UmsMenu {
	var umsMenu model.UmsMenu
	result := iService.DB.First(&umsMenu, id)
	if result.Error != nil {
		return nil
	}
	return &umsMenu
}

// 删除菜单
func (iService UmsMenuService) Delete(id string) error {
	result := iService.DB.Delete(&model.UmsMenu{}, id)
	return result.Error
}

// 获取菜单列表
func (iService UmsMenuService) List(parentIdStr, pageStr, sizeStr string) ([]model.UmsMenu, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	where := iService.DB.Where("parent_id = ?", parentIdStr)
	var count int64
	result := where.Model(&model.UmsMenu{}).Count(&count)
	if result.Error != nil || count == 0 {
		return nil, 0
	}
	var list []model.UmsMenu
	result = where.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

// 根据用户id获取菜单
func (iService UmsMenuService) GetMenuList(adminId string) []model.UmsMenu {
	var list []model.UmsMenu
	result := iService.DB.Raw("SELECT m.id id, m.parent_id, m.create_time, m.title, m.level, m.sort, m.name, m.icon, m.hidden FROM ums_admin_role_relation arr LEFT JOIN ums_role_menu_relation rmr ON arr.role_id = rmr.role_id LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE arr.admin_id = ? AND m.id IS NOT NULL GROUP BY m.id", adminId).Scan(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

// 获取树形菜单
func (iService UmsMenuService) TreeList() []*dto.UmsMenuDto {
	var list []model.UmsMenu
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	res := make([]*dto.UmsMenuDto, 0)
	menuMap := make(map[int64]*dto.UmsMenuDto)
	for _, menu := range list {
		menuDto := dto.UmsMenuDto{
			UmsMenu:  menu,
			Children: make([]dto.UmsMenuDto, 0),
		}
		menuMap[menu.ID] = &menuDto
		if menu.ParentID == 0 {
			res = append(res, &menuDto)
		} else {
			if parent, exist := menuMap[menu.ParentID]; exist {
				parent.Children = append(parent.Children, menuDto)
			}
		}
	}
	return res
}

// 修改显示状态
func (iService UmsMenuService) UpdateHidden(id, hidden string) error {
	var umsMenu model.UmsMenu
	result := iService.DB.Model(&umsMenu).Where("id = ?", id).Update("hidden", hidden)
	return result.Error
}
