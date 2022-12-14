// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductAttribute = "pms_product_attribute"

// PmsProductAttribute mapped from table <pms_product_attribute>
type PmsProductAttribute struct {
	ID                         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	ProductAttributeCategoryID int64  `gorm:"column:product_attribute_category_id" json:"productAttributeCategoryId" form:"productAttributeCategoryId"`
	Name                       string `gorm:"column:name" json:"name" form:"name"`
	SelectType                 int32  `gorm:"column:select_type" json:"selectType" form:"selectType"`           // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int32  `gorm:"column:input_type" json:"inputType" form:"inputType"`              // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string `gorm:"column:input_list" json:"inputList" form:"inputList"`              // 可选值列表，以逗号隔开
	Sort                       int32  `gorm:"column:sort" json:"sort" form:"sort"`                              // 排序字段：最高的可以单独上传图片
	FilterType                 int32  `gorm:"column:filter_type" json:"filterType" form:"filterType"`           // 分类筛选样式：1->普通；1->颜色
	SearchType                 int32  `gorm:"column:search_type" json:"searchType" form:"searchType"`           // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int32  `gorm:"column:related_status" json:"relatedStatus" form:"relatedStatus"`  // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              int32  `gorm:"column:hand_add_status" json:"handAddStatus" form:"handAddStatus"` // 是否支持手动新增；0->不支持；1->支持
	Type                       int32  `gorm:"column:type" json:"type" form:"type"`                              // 属性的类型；0->规格；1->参数
}

// TableName PmsProductAttribute's table name
func (*PmsProductAttribute) TableName() string {
	return TableNamePmsProductAttribute
}
