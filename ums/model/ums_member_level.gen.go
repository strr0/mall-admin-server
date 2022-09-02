// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUmsMemberLevel = "ums_member_level"

// UmsMemberLevel mapped from table <ums_member_level>
type UmsMemberLevel struct {
	ID                    int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	Name                  string  `gorm:"column:name" json:"name" form:"name"`
	GrowthPoint           int32   `gorm:"column:growth_point" json:"growthPoint" form:"growthPoint"`
	DefaultStatus         int32   `gorm:"column:default_status" json:"defaultStatus" form:"defaultStatus"`                          // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      float64 `gorm:"column:free_freight_point" json:"freeFreightPoint" form:"freeFreightPoint"`                // 免运费标准
	CommentGrowthPoint    int32   `gorm:"column:comment_growth_point" json:"commentGrowthPoint" form:"commentGrowthPoint"`          // 每次评价获取的成长值
	PriviledgeFreeFreight int32   `gorm:"column:priviledge_free_freight" json:"priviledgeFreeFreight" form:"priviledgeFreeFreight"` // 是否有免邮特权
	PriviledgeSignIn      int32   `gorm:"column:priviledge_sign_in" json:"priviledgeSignIn" form:"priviledgeSignIn"`                // 是否有签到特权
	PriviledgeComment     int32   `gorm:"column:priviledge_comment" json:"priviledgeComment" form:"priviledgeComment"`              // 是否有评论获奖励特权
	PriviledgePromotion   int32   `gorm:"column:priviledge_promotion" json:"priviledgePromotion" form:"priviledgePromotion"`        // 是否有专享活动特权
	PriviledgeMemberPrice int32   `gorm:"column:priviledge_member_price" json:"priviledgeMemberPrice" form:"priviledgeMemberPrice"` // 是否有会员价格特权
	PriviledgeBirthday    int32   `gorm:"column:priviledge_birthday" json:"priviledgeBirthday" form:"priviledgeBirthday"`           // 是否有生日特权
	Note                  string  `gorm:"column:note" json:"note" form:"note"`
}

// TableName UmsMemberLevel's table name
func (*UmsMemberLevel) TableName() string {
	return TableNameUmsMemberLevel
}
