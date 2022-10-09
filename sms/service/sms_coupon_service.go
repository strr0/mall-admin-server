package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
)

// 优惠券管理
type SmsCouponService struct {
	DB *gorm.DB
}

// 新增优惠券
func (iService SmsCouponService) Create(smsCouponDto dto.SmsCouponDto) error {
	smsCoupon := smsCouponDto.SmsCoupon
	smsCoupon.Count_ = smsCoupon.PublishCount
	smsCoupon.UseCount = 0
	smsCoupon.ReceiveCount = 0
	result := iService.DB.Create(&smsCoupon)
	if result.Error != nil {
		return result.Error
	}
	return iService.UpdateRelation(smsCouponDto)
}

// 删除优惠券
func (iService SmsCouponService) Delete(id string) error {
	result := iService.DB.Delete(&model.SmsCoupon{}, id)
	if result.Error != nil {
		return result.Error
	}
	return iService.DeleteRelation(id)
}

// 删除关联
func (iService SmsCouponService) DeleteRelation(id string) error {
	result := iService.DB.Where("coupon_id = ?", id).Delete(&model.SmsCouponProductCategoryRelation{})
	if result.Error != nil {
		return result.Error
	}
	result = iService.DB.Where("coupon_id = ?", id).Delete(&model.SmsCouponProductRelation{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新优惠券
func (iService SmsCouponService) Update(id string, smsCouponDto dto.SmsCouponDto) error {
	result := iService.DB.Save(&smsCouponDto.SmsCoupon)
	if result.Error != nil {
		return result.Error
	}
	err := iService.DeleteRelation(id)
	if err != nil {
		return err
	}
	return iService.UpdateRelation(smsCouponDto)
}

// 更新关联
func (iService SmsCouponService) UpdateRelation(smsCouponDto dto.SmsCouponDto) error {
	var result *gorm.DB
	switch smsCouponDto.UseType {
	case 1:
		for _, productCategoryRelation := range smsCouponDto.ProductCategoryRelationList {
			productCategoryRelation.CouponID = smsCouponDto.ID
			result = iService.DB.Create(&productCategoryRelation)
		}
	case 2:
		for _, productRelation := range smsCouponDto.ProductRelationList {
			productRelation.CouponID = smsCouponDto.ID
			result = iService.DB.Create(&productRelation)
		}
	default:
	}
	return result.Error
}

// 优惠券列表
func (iService SmsCouponService) List(name, type_, pageStr, sizeStr string) ([]model.SmsCoupon, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if name != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("name like ?", "%"+name+"%")
		})
	}
	if type_ != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("type = ?", type_)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsCoupon{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsCoupon
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}

// 优惠券详情
func (iService SmsCouponService) GetItem(id string) *model.SmsCoupon {
	var smsCoupon model.SmsCoupon
	result := iService.DB.First(&smsCoupon, id)
	if result.Error != nil {
		return nil
	}
	return &smsCoupon
}
