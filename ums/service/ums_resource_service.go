package service

import (
	"gorm.io/gen"
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/util"
)

type UmsResourceService struct {
	//
}

// 创建资源
func (UmsResourceService) Create(umsResource model.UmsResource) error {
	return query.UmsResource.Create(&umsResource)
}

// 更新资源
func (UmsResourceService) Update(idStr string, umsResource model.UmsResource) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsResource.Where(query.UmsResource.ID.Eq(id)).Updates(umsResource)
	return err
}

// 根据id获取资源
func (UmsResourceService) GetItem(idStr string) *model.UmsResource {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.UmsResource.Where(query.UmsResource.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

// 删除资源
func (UmsResourceService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsResource.Where(query.UmsResource.ID.Eq(id)).Delete()
	return err
}

// 资源列表
func (UmsResourceService) List(categoryIdStr, nameKeyword, urlKeyword, pageStr, sizeStr string) ([]*model.UmsResource, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	umsResource := query.UmsResource
	categoryId, err := util.ParseInt64WithErr(categoryIdStr)
	conds := make([]gen.Condition, 0)
	if err == nil {
		conds = append(conds, umsResource.CategoryID.Eq(categoryId))
	}
	if nameKeyword != "" {
		conds = append(conds, umsResource.Name.Like(nameKeyword))
	}
	if urlKeyword != "" {
		conds = append(conds, umsResource.URL.Like(urlKeyword))
	}
	umsResource.Where(conds...)
	total, err := umsResource.Count()
	if err != nil || total == 0 {
		return nil, 0
	}
	find, err := umsResource.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

// 获取全部资源
func (UmsResourceService) ListAll() []*model.UmsResource {
	find, err := query.UmsResource.Find()
	if err != nil {
		return nil
	}
	return find
}

// 根据用户id获取资源
func (UmsResourceService) GetResourceList(adminId int64) []*model.UmsResource {
	list, err := query.UmsResource.GetResourceList(adminId)
	if err != nil {
		return nil
	}
	return list
}
