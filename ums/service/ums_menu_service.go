package service

import (
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/ums/service/dto"
	"mall-admin-server/util"
	"time"
)

type UmsMenuService struct {
	//
}

// 创建菜单
func (UmsMenuService) Create(umsMenu model.UmsMenu) error {
	umsMenu.CreateTime = time.Now()
	return query.UmsMenu.Create(&umsMenu)
}

// 更新菜单
func (UmsMenuService) Update(idStr string, umsMenu model.UmsMenu) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsMenu.Where(query.UmsMenu.ID.Eq(id)).Updates(umsMenu)
	return err
}

// 根据id获取菜单
func (UmsMenuService) GetItem(idStr string) *model.UmsMenu {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.UmsMenu.Where(query.UmsMenu.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

// 删除菜单
func (UmsMenuService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsMenu.Where(query.UmsMenu.ID.Eq(id)).Delete()
	return err
}

// 获取菜单列表
func (UmsMenuService) List(parentIdStr, pageStr, sizeStr string) ([]*model.UmsMenu, int64) {
	parentId, err := util.ParseInt64WithErr(parentIdStr)
	if err != nil {
		return nil, 0
	}
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	umsMenu := query.UmsMenu.Where(query.UmsMenu.ParentID.Eq(parentId))
	total, err := umsMenu.Count()
	if err != nil || total == 0 {
		return nil, 0
	}
	find, err := umsMenu.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

// 根据用户id获取菜单
func (UmsMenuService) GetMenuList(adminIdStr string) []*model.UmsMenu {
	adminId, err := util.ParseInt64WithErr(adminIdStr)
	if err != nil {
		return nil
	}
	list, err := query.UmsMenu.GetMenuList(adminId)
	if err != nil {
		return nil
	}
	return list
}

// 获取树形菜单
func (UmsMenuService) TreeList() []*dto.UmsMenuDto {
	find, err := query.UmsMenu.Find()
	if err != nil {
		return nil
	}
	res := make([]*dto.UmsMenuDto, 0)
	menuMap := make(map[int64]*dto.UmsMenuDto)
	for _, menu := range find {
		menuDto := dto.UmsMenuDto{
			UmsMenu:  *menu,
			Children: make([]*dto.UmsMenuDto, 0),
		}
		menuMap[menu.ID] = &menuDto
		if menu.ParentID == 0 {
			res = append(res, &menuDto)
		} else {
			if parent, exist := menuMap[menu.ParentID]; exist {
				parent.Children = append(parent.Children, &menuDto)
			}
		}
	}
	return res
}

// 修改显示状态
func (UmsMenuService) UpdateHidden(idStr, hiddenStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	hidden, err := util.ParseInt32WithErr(hiddenStr)
	if err != nil {
		return err
	}
	var umsMenu model.UmsMenu
	umsMenu.Hidden = hidden
	_, err = query.UmsMenu.Where(query.UmsMenu.ID.Eq(id)).Updates(umsMenu)
	return err
}
