// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCmsSubjectProductRelation = "cms_subject_product_relation"

// CmsSubjectProductRelation mapped from table <cms_subject_product_relation>
type CmsSubjectProductRelation struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	SubjectID int64 `gorm:"column:subject_id" json:"subjectId" form:"subjectId"`
	ProductID int64 `gorm:"column:product_id" json:"productId" form:"productId"`
}

// TableName CmsSubjectProductRelation's table name
func (*CmsSubjectProductRelation) TableName() string {
	return TableNameCmsSubjectProductRelation
}