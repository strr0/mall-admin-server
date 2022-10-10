package service

import (
	"gorm.io/gorm"
	"mall-admin-server/cms/model"
	"mall-admin-server/util"
)

// 商品专题管理
type CmsSubjectService struct {
	DB *gorm.DB
}

func (iService CmsSubjectService) ListAll() []model.CmsSubject {
	var list []model.CmsSubject
	result := iService.DB.Find(&list)
	if result.Error != nil {
		return nil
	}
	return list
}

func (iService CmsSubjectService) List(keyword, pageStr, sizeStr string) ([]model.CmsSubject, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if keyword != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("title like ?", "%"+keyword+"%")
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.CmsSubject{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.CmsSubject
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
