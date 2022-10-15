package model

import "time"

type MemberReadHistory struct {
	Id              string    `json:"id" form:"id"`
	MemberId        string    `json:"memberId" form:"memberId"`
	MemberNickName  string    `json:"memberNickName" form:"memberNickName"`
	MemberIcon      string    `json:"memberIcon" form:"memberIcon"`
	ProductId       string    `json:"productId" form:"productId"`
	ProductName     string    `json:"productName" form:"productName"`
	ProductPic      string    `json:"productPic" form:"productPic"`
	ProductSubTitle string    `json:"productSubTitle" form:"productSubTitle"`
	ProductPrice    string    `json:"productPrice" form:"productPrice"`
	CreateDate      time.Time `json:"createDate" form:"createDate"`
}
