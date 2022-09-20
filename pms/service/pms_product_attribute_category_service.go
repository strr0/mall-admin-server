package service

import (
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/query"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品属性分类管理
type PmsProductAttributeCategoryService struct {
	//
}

func (PmsProductAttributeCategoryService) Create(name string) error {
	var pmsProductAttributeCategory model.PmsProductAttributeCategory
	pmsProductAttributeCategory.Name = name
	return query.PmsProductAttributeCategory.Create(&pmsProductAttributeCategory)
}

func (PmsProductAttributeCategoryService) Update(idStr, name string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	var pmsProductAttributeCategory model.PmsProductAttributeCategory
	pmsProductAttributeCategory.Name = name
	_, err = query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(id)).Updates(pmsProductAttributeCategory)
	return err
}

func (PmsProductAttributeCategoryService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(id)).Delete()
	return err
}

func (PmsProductAttributeCategoryService) GetItem(idStr string) *model.PmsProductAttributeCategory {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.PmsProductAttributeCategory.Where(query.PmsProductAttributeCategory.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (PmsProductAttributeCategoryService) List(pageStr, sizeStr string) ([]*model.PmsProductAttributeCategory, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	pmsProductAttributeCategory := query.PmsProductAttributeCategory
	total, err := pmsProductAttributeCategory.Count()
	if err != nil {
		return nil, 0
	}
	find, err := pmsProductAttributeCategory.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

func (PmsProductAttributeCategoryService) GetListWithAttr() []*dto.PmsProductAttributeCategoryDto {
	attr, err := query.PmsProductAttributeCategory.GetListWithAttr()
	if err != nil {
		return nil
	}
	return attr
}
