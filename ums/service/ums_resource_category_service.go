package service

import (
	"mall-admin-server/ums/model"
	"mall-admin-server/ums/query"
	"mall-admin-server/util"
)

type UmsResourceCategoryService struct {
	//
}

// 获取所有分类
func (UmsResourceCategoryService) ListAll() []*model.UmsResourceCategory {
	find, err := query.UmsResourceCategory.Find()
	if err != nil {
		return nil
	}
	return find
}

// 创建分类
func (UmsResourceCategoryService) Create(umsResourceCategory model.UmsResourceCategory) error {
	return query.UmsResourceCategory.Create(&umsResourceCategory)
}

// 修改分类
func (UmsResourceCategoryService) Update(idStr string, umsResourceCategory model.UmsResourceCategory) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsResourceCategory.Where(query.UmsResourceCategory.ID.Eq(id)).Updates(umsResourceCategory)
	return err
}

// 删除分类
func (UmsResourceCategoryService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.UmsResourceCategory.Where(query.UmsResourceCategory.ID.Eq(id)).Delete()
	return err
}
