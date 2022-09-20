package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/util"
)

type SmsCouponHistoryService struct {
	//
}

// 优惠券领取记录
func (SmsCouponHistoryService) List(couponIdStr, useStatusStr, orderSn, pageStr, sizeStr string) ([]*model.SmsCouponHistory, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsCouponHistory := query.SmsCouponHistory
	conds := make([]gen.Condition, 0)
	if couponId, err := util.ParseInt64WithErr(couponIdStr); err == nil {
		conds = append(conds, smsCouponHistory.CouponID.Eq(couponId))
	}
	if useStatus, err := util.ParseInt32WithErr(useStatusStr); err == nil {
		conds = append(conds, smsCouponHistory.UseStatus.Eq(useStatus))
	}
	if orderSn != "" {
		conds = append(conds, smsCouponHistory.OrderSn.Eq(orderSn))
	}
	smsCouponHistory.Where(conds...)
	total, err := smsCouponHistory.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsCouponHistory.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}
