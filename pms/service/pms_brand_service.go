package service

import (
	"gorm.io/gorm"
	"mall-admin-server/pms/model"
	"mall-admin-server/util"
)

// 商品品牌管理
type PmsBrandService struct {
	DB *gorm.DB
}

func (iService PmsBrandService) ListAll() []model.PmsBrand {
	var list []model.PmsBrand
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService PmsBrandService) Create(pmsBrand model.PmsBrand) error {
	if pmsBrand.FirstLetter == "" {
		pmsBrand.FirstLetter = string(pmsBrand.Name[0])
	}
	result := iService.DB.Create(&pmsBrand)
	return result.Error
}

func (iService PmsBrandService) Update(id string, pmsBrand model.PmsBrand) error {
	if pmsBrand.FirstLetter == "" {
		pmsBrand.FirstLetter = string(pmsBrand.Name[0])
	}
	result := iService.DB.Save(&pmsBrand)
	if result.Error != nil {
		return result.Error
	}
	result = iService.DB.Model(&model.PmsProduct{}).Where("brand_id = ?", id).Update("brand_name", pmsBrand.Name)
	return result.Error
}

func (iService PmsBrandService) Delete(id string) error {
	result := iService.DB.Delete(&model.PmsBrand{}, id)
	return result.Error
}

func (iService PmsBrandService) DeleteBrand(ids []string) error {
	result := iService.DB.Delete(&model.PmsBrand{}, ids)
	return result.Error
}

func (iService PmsBrandService) List(keyword, pageStr, sizeStr string) ([]model.PmsBrand, int64) {
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
	result := scopes.Model(&model.PmsBrand{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.PmsBrand
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

func (iService PmsBrandService) GetItem(id string) *model.PmsBrand {
	var pmsBrand model.PmsBrand
	result := iService.DB.First(&pmsBrand, id)
	if result.Error != nil {
		return nil
	}
	return &pmsBrand
}

func (iService PmsBrandService) UpdateShowStatus(ids []string, showStatus string) error {
	result := iService.DB.Model(&model.PmsBrand{}).Where("id in ?", ids).Update("show_status", showStatus)
	return result.Error
}

func (iService PmsBrandService) UpdateFactoryStatus(ids []string, factoryStatus string) error {
	result := iService.DB.Model(&model.PmsBrand{}).Where("id in ?", ids).Update("factory_status", factoryStatus)
	return result.Error
}
