// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSmsCoupon = "sms_coupon"

// SmsCoupon mapped from table <sms_coupon>
type SmsCoupon struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	Type         int32     `gorm:"column:type" json:"type" form:"type"` // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         string    `gorm:"column:name" json:"name" form:"name"`
	Platform     int32     `gorm:"column:platform" json:"platform" form:"platform"`  // 使用平台：0->全部；1->移动；2->PC
	Count_       int32     `gorm:"column:count" json:"count" form:"count"`           // 数量
	Amount       float64   `gorm:"column:amount" json:"amount" form:"amount"`        // 金额
	PerLimit     int32     `gorm:"column:per_limit" json:"perLimit" form:"perLimit"` // 每人限领张数
	MinPoint     float64   `gorm:"column:min_point" json:"minPoint" form:"minPoint"` // 使用门槛；0表示无门槛
	StartTime    time.Time `gorm:"column:start_time" json:"startTime" form:"startTime"`
	EndTime      time.Time `gorm:"column:end_time" json:"endTime" form:"endTime"`
	UseType      int32     `gorm:"column:use_type" json:"useType" form:"useType"`                // 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         string    `gorm:"column:note" json:"note" form:"note"`                          // 备注
	PublishCount int32     `gorm:"column:publish_count" json:"publishCount" form:"publishCount"` // 发行数量
	UseCount     int32     `gorm:"column:use_count" json:"useCount" form:"useCount"`             // 已使用数量
	ReceiveCount int32     `gorm:"column:receive_count" json:"receiveCount" form:"receiveCount"` // 领取数量
	EnableTime   time.Time `gorm:"column:enable_time" json:"enableTime" form:"enableTime"`       // 可以领取的日期
	Code         string    `gorm:"column:code" json:"code" form:"code"`                          // 优惠码
	MemberLevel  int32     `gorm:"column:member_level" json:"memberLevel" form:"memberLevel"`    // 可领取的会员类型：0->无限时
}

// TableName SmsCoupon's table name
func (*SmsCoupon) TableName() string {
	return TableNameSmsCoupon
}