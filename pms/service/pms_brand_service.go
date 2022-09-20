package service

import (
	"gorm.io/gen"
	"mall-admin-server/pms/model"
	"mall-admin-server/pms/query"
	"mall-admin-server/util"
)

// 商品品牌管理
type PmsBrandService struct {
	//
}

func (PmsBrandService) ListAll() []*model.PmsBrand {
	find, err := query.PmsBrand.Find()
	if err != nil {
		return nil
	}
	return find
}

func (PmsBrandService) Create(pmsBrand model.PmsBrand) error {
	if pmsBrand.FirstLetter == "" {
		pmsBrand.FirstLetter = string(pmsBrand.Name[0])
	}
	return query.PmsBrand.Create(&pmsBrand)
}

func (PmsBrandService) Update(idStr string, pmsBrand model.PmsBrand) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	if pmsBrand.FirstLetter == "" {
		pmsBrand.FirstLetter = string(pmsBrand.Name[0])
	}
	_, err = query.PmsBrand.Where(query.PmsBrand.ID.Eq(id)).Updates(pmsBrand)
	if err != nil {
		return err
	}
	var pmsProduct model.PmsProduct
	pmsProduct.BrandName = pmsBrand.Name
	_, err = query.PmsProduct.Where(query.PmsProduct.BrandID.Eq(id)).Updates(pmsProduct)
	return err
}

func (PmsBrandService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.PmsBrand.Where(query.PmsBrand.ID.Eq(id)).Delete()
	return err
}

func (PmsBrandService) DeleteBrand(idsStr []string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	_, err := query.PmsBrand.Where(query.PmsBrand.ID.In(ids...)).Delete()
	return err
}

func (PmsBrandService) List(keyword, pageStr, sizeStr string) ([]*model.PmsBrand, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	pmsBrand := query.PmsBrand
	conds := make([]gen.Condition, 0)
	if keyword != "" {
		conds = append(conds, query.PmsBrand.Name.Like("%"+keyword+"%"))
	}
	pmsBrandDo := pmsBrand.Where(conds...)
	total, err := pmsBrandDo.Count()
	if err != nil {
		return nil, 0
	}
	find, err := pmsBrandDo.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

func (PmsBrandService) GetItem(idStr string) *model.PmsBrand {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.PmsBrand.Where(query.PmsBrand.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}

func (PmsBrandService) UpdateShowStatus(idsStr []string, showStatusStr string) error {
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
	var pmsBrand model.PmsBrand
	pmsBrand.ShowStatus = showStatus
	_, err = query.PmsBrand.Where(query.PmsBrand.ID.In(ids...)).Updates(pmsBrand)
	return err
}

func (PmsBrandService) UpdateFactoryStatus(idsStr []string, factoryStatusStr string) error {
	ids := make([]int64, 0)
	for _, idStr := range idsStr {
		id, err := util.ParseInt64WithErr(idStr)
		if err == nil {
			ids = append(ids, id)
		}
	}
	factoryStatus, err := util.ParseInt32WithErr(factoryStatusStr)
	if err != nil {
		return err
	}
	var pmsBrand model.PmsBrand
	pmsBrand.FactoryStatus = factoryStatus
	_, err = query.PmsBrand.Where(query.PmsBrand.ID.In(ids...)).Updates(pmsBrand)
	return err
}
