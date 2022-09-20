// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOmsOrderOperateHistory = "oms_order_operate_history"

// OmsOrderOperateHistory mapped from table <oms_order_operate_history>
type OmsOrderOperateHistory struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	OrderID     int64     `gorm:"column:order_id" json:"orderId" form:"orderId"`             // 订单id
	OperateMan  string    `gorm:"column:operate_man" json:"operateMan" form:"operateMan"`    // 操作人：用户；系统；后台管理员
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime" form:"createTime"`    // 操作时间
	OrderStatus int32     `gorm:"column:order_status" json:"orderStatus" form:"orderStatus"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    `gorm:"column:note" json:"note" form:"note"`                       // 备注
}

// TableName OmsOrderOperateHistory's table name
func (*OmsOrderOperateHistory) TableName() string {
	return TableNameOmsOrderOperateHistory
}