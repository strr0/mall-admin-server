package service

import (
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/query"
	"mall-admin-server/pms/service/dto"
	"mall-admin-server/util"
)

// 商品分类管理
type PmsProductCategoryService struct {
	//
}

func (iService PmsProductCategoryService) Create(pmsProductCategoryDto dto.PmsProductCategoryDto) error {
	pmsProductCategory := pmsProductCategoryDto.PmsProductCategory
	pmsProductCategory.ProductCount = 0
	iService.setCategoryLevel(&pmsProductCategory)
	err := query.PmsProductCategory.Create(&pmsProductCategory)
	if err != nil {
		return err
	}
	// 创建关联属性
	if len(pmsProductCategoryDto.ProductAttributeIdList) > 0 {
		err := iService.insertRelationList(pmsProductCategoryDto.ID, pmsProductCategoryDto.ProductAttributeIdList)
		if err != nil {
			return err
		}
	}
	return nil
}

func (iService PmsProductCategoryService) Update(idStr string, pmsProductCategoryDto dto.PmsProductCategoryDto) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	pmsProductCategory := pmsProductCategoryDto.PmsProductCategory
	iService.setCategoryLevel(&pmsProductCategory)
	// 更新商品的分类名称
	var pmsProduct model.PmsProduct
	pmsProduct.ProductCategoryName = pmsProductCategory.Name
	_, err = query.PmsProduct.Where(query.PmsProduct.ProductCategoryID.Eq(id)).Updates(pmsProduct)
	if err != nil {
		return err
	}
	// 更新关联属性
	_, err = query.PmsProductCategoryAttributeRelation.Where(query.PmsProductCategoryAttributeRelation.ProductCategoryID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	if len(pmsProductCategoryDto.ProductAttributeIdList) > 0 {
		err = iService.insertRelationList(id, pmsProductCategoryDto.ProductAttributeIdList)
		if err != nil {
			return err
		}
	}
	_, err = query.PmsProductCategory.Where(query.PmsProductCategory.ID.Eq(id)).Updates(pmsProductCategory)
	return err
}

func (PmsProductCategoryService) List(parentIdStr, pageStr, sizeStr string) ([]*model.PmsProductCategory, int64) {
	parentId, err := util.ParseInt64WithErr(parentIdStr)
	if err != nil {
		return nil, 0
	}
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	pmsProductCategoryDo := query.PmsProductCategory.Where(query.PmsProductCategory.ParentID.Eq(parentId))
	total, err := pmsProductCategoryDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := pmsProductCategoryDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

func (PmsProductCategoryService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.PmsProductCategory.Where(query.PmsProductCategory.ID.Eq(id)).Delete()
	return err
}

func (PmsProductCategoryService) GetItem(idStr string) *model.PmsProductCategory {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.PmsProductCategory.Where(query.PmsProductCategory.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (PmsProductCategoryService) UpdateNavStatus(idsStr []string, navStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	navStatus, err := util.ParseInt32WithErr(navStatusStr)
	if err != nil {
		return err
	}
	var pmsProductCategory model.PmsProductCategory
	pmsProductCategory.NavStatus = navStatus
	_, err = query.PmsProductCategory.Where(query.PmsProductCategory.ID.In(ids...)).Updates(pmsProductCategory)
	return err
}

func (PmsProductCategoryService) UpdateShowStatus(idsStr []string, showStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	showStatus, err := util.ParseInt32WithErr(showStatusStr)
	if err != nil {
		return err
	}
	var pmsProductCategory model.PmsProductCategory
	pmsProductCategory.ShowStatus = showStatus
	_, err = query.PmsProductCategory.Where(query.PmsProductCategory.ID.In(ids...)).Updates(pmsProductCategory)
	return err
}

func (PmsProductCategoryService) TreeList() []*dto.PmsProductCategoryTreeDto {
	find, err := query.PmsProductCategory.Find()
	if err != nil {
		return nil
	}
	categoryMap := make(map[int64]*dto.PmsProductCategoryTreeDto)
	res := make([]*dto.PmsProductCategoryTreeDto, 0)
	for _, category := range find {
		categoryDto := dto.PmsProductCategoryTreeDto{
			PmsProductCategory: *category,
			Children:           make([]dto.PmsProductCategoryTreeDto, 0),
		}
		categoryMap[category.ID] = &categoryDto
		if category.ParentID == 0 {
			res = append(res, &categoryDto)
		} else {
			if parent, exist := categoryMap[category.ParentID]; exist {
				parent.Children = append(parent.Children, categoryDto)
			}
		}
	}
	return res
}

// 创建关联属性
func (PmsProductCategoryService) insertRelationList(productCategoryId int64, productAttributeIdList []int64) error {
	values := make([]*model.PmsProductCategoryAttributeRelation, 0)
	for _, productAttributeId := range productAttributeIdList {
		values = append(values, &model.PmsProductCategoryAttributeRelation{
			ProductAttributeID: productAttributeId,
			ProductCategoryID:  productCategoryId,
		})
	}
	return query.PmsProductCategoryAttributeRelation.Create(values...)
}

// 设置分类level
func (PmsProductCategoryService) setCategoryLevel(pmsProductCategory *model.PmsProductCategory) {
	if pmsProductCategory.ParentID == 0 {
		pmsProductCategory.Level = 0
	} else {
		first, err := query.PmsProductCategory.Where(query.PmsProductCategory.ID.Eq(pmsProductCategory.ParentID)).First()
		if err != nil {
			return
		}
		if first != nil {
			pmsProductCategory.Level = first.Level + 1
		} else {
			pmsProductCategory.Level = 0
		}
	}
}
