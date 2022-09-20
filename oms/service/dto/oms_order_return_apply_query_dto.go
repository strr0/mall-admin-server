package dto

import "time"

type OmsOrderReturnApplyQueryDto struct {
	Id              int64     `json:"id" form:"id"`
	ReceiverKeyword string    `json:"receiverKeyword" form:"receiverKeyword"`
	Status          int32     `json:"status" form:"status"`
	CreateTime      time.Time `json:"createTime" form:"createTime"`
	HandleMan       string    `json:"handleMan" form:"handleMan"`
	HandleTime      time.Time `json:"handleTime" form:"handleTime"`
}
