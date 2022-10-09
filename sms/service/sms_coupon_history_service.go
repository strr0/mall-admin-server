package service

import (
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/util"
)

type SmsCouponHistoryService struct {
	DB *gorm.DB
}

// 优惠券领取记录
func (iService SmsCouponHistoryService) List(couponId, useStatus, orderSn, pageStr, sizeStr string) ([]model.SmsCouponHistory, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	funcs := make([]func(*gorm.DB) *gorm.DB, 0)
	if couponId != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("coupon_id = ?", couponId)
		})
	}
	if useStatus != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("use_status = ?", useStatus)
		})
	}
	if orderSn != "" {
		funcs = append(funcs, func(db *gorm.DB) *gorm.DB {
			return db.Where("order_sn = ?", orderSn)
		})
	}
	scopes := iService.DB.Scopes(funcs...)
	var count int64
	result := scopes.Model(&model.SmsCouponHistory{}).Count(&count)
	if result.Error != nil {
		return nil, 0
	}
	var list []model.SmsCouponHistory
	result = scopes.Offset(offset).Limit(size).Find(&list)
	if result.Error != nil {
		return nil, count
	}
	return list, count
}
