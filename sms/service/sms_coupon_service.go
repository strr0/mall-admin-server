package service

import (
	"gorm.io/gen"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/query"
	"mall-admin-server/sms/service/dto"
	"mall-admin-server/util"
)

// 优惠券管理
type SmsCouponService struct {
	//
}

// 新增优惠券
func (iService SmsCouponService) Create(smsCouponDto dto.SmsCouponDto) error {
	smsCoupon := smsCouponDto.SmsCoupon
	smsCoupon.Count_ = smsCoupon.PublishCount
	smsCoupon.UseCount = 0
	smsCoupon.ReceiveCount = 0
	err := query.SmsCoupon.Create(&smsCoupon)
	if err != nil {
		return err
	}
	return iService.UpdateRelation(smsCouponDto)
}

// 删除优惠券
func (iService SmsCouponService) Delete(idStr string) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsCoupon.Where(query.SmsCoupon.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return iService.DeleteRelation(id)
}

// 删除关联
func (SmsCouponService) DeleteRelation(id int64) error {
	_, err := query.SmsCouponProductCategoryRelation.Where(query.SmsCouponProductCategoryRelation.CouponID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	_, err = query.SmsCouponProductRelation.Where(query.SmsCouponProductRelation.CouponID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// 更新优惠券
func (iService SmsCouponService) Update(idStr string, smsCouponDto dto.SmsCouponDto) error {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return err
	}
	_, err = query.SmsCoupon.Where(query.SmsCoupon.ID.Eq(id)).Updates(smsCouponDto.SmsCoupon)
	if err != nil {
		return err
	}
	err = iService.DeleteRelation(id)
	if err != nil {
		return err
	}
	return iService.UpdateRelation(smsCouponDto)
}

// 更新关联
func (SmsCouponService) UpdateRelation(smsCouponDto dto.SmsCouponDto) error {
	var err error
	switch smsCouponDto.UseType {
	case 1:
		for _, productCategoryRelation := range smsCouponDto.ProductCategoryRelationList {
			productCategoryRelation.CouponID = smsCouponDto.ID
			err = query.SmsCouponProductCategoryRelation.Create(&productCategoryRelation)
		}
	case 2:
		for _, productRelation := range smsCouponDto.ProductRelationList {
			productRelation.CouponID = smsCouponDto.ID
			err = query.SmsCouponProductRelation.Create(&productRelation)
		}
	default:
	}
	return err
}

// 优惠券列表
func (SmsCouponService) List(name, typeStr, pageStr, sizeStr string) ([]*model.SmsCoupon, int64) {
	page := util.ParseInt(pageStr, 1)
	size := util.ParseInt(sizeStr, 10)
	offset := (page - 1) * size
	smsCoupon := query.SmsCoupon
	conds := make([]gen.Condition, 0)
	if name != "" {
		conds = append(conds, smsCoupon.Name.Like("%"+name+"%"))
	}
	if type_, err := util.ParseInt32WithErr(typeStr); err != nil {
		conds = append(conds, smsCoupon.Type.Eq(type_))
	}
	smsCoupon.Where(conds...)
	total, err := smsCoupon.Count()
	if err != nil {
		return nil, 0
	}
	find, err := smsCoupon.Offset(offset).Limit(size).Find()
	if err != nil {
		return nil, total
	}
	return find, total
}

// 优惠券详情
func (SmsCouponService) GetItem(idStr string) *model.SmsCoupon {
	id, err := util.ParseInt64WithErr(idStr)
	if err != nil {
		return nil
	}
	first, err := query.SmsCoupon.Where(query.SmsCoupon.ID.Eq(id)).First()
	if err != nil {
		return nil
	}
	return first
}
