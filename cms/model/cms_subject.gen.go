// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCmsSubject = "cms_subject"

// CmsSubject mapped from table <cms_subject>
type CmsSubject struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id" form:"id"`
	CategoryID      int64     `gorm:"column:category_id" json:"categoryId" form:"categoryId"`
	Title           string    `gorm:"column:title" json:"title" form:"title"`
	Pic             string    `gorm:"column:pic" json:"pic" form:"pic"`                             // 专题主图
	ProductCount    int32     `gorm:"column:product_count" json:"productCount" form:"productCount"` // 关联产品数量
	RecommendStatus int32     `gorm:"column:recommend_status" json:"recommendStatus" form:"recommendStatus"`
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime" form:"createTime"`
	CollectCount    int32     `gorm:"column:collect_count" json:"collectCount" form:"collectCount"`
	ReadCount       int32     `gorm:"column:read_count" json:"readCount" form:"readCount"`
	CommentCount    int32     `gorm:"column:comment_count" json:"commentCount" form:"commentCount"`
	AlbumPics       string    `gorm:"column:album_pics" json:"albumPics" form:"albumPics"` // 画册图片用逗号分割
	Description     string    `gorm:"column:description" json:"description" form:"description"`
	ShowStatus      int32     `gorm:"column:show_status" json:"showStatus" form:"showStatus"` // 显示状态：0->不显示；1->显示
	Content         string    `gorm:"column:content" json:"content" form:"content"`
	ForwardCount    int32     `gorm:"column:forward_count" json:"forwardCount" form:"forwardCount"` // 转发数
	CategoryName    string    `gorm:"column:category_name" json:"categoryName" form:"categoryName"` // 专题分类名称
}

// TableName CmsSubject's table name
func (*CmsSubject) TableName() string {
	return TableNameCmsSubject
}
