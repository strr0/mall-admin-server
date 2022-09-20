package cmd

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"mall-admin-server/sms/model"
	"mall-admin-server/sms/service/dto"
)

type SmsFlashPromotionProductRelationMethod interface {
	// sql(SELECT COUNT(1) from (SELECT r.id, r.flash_promotion_price, r.flash_promotion_count, r.flash_promotion_limit,
	//            r.flash_promotion_id, r.flash_promotion_session_id, r.product_id, r.sort,
	//            p.`name` name, p.product_sn product_sn, p.price price, p.stock stock
	// FROM sms_flash_promotion_product_relation r
	// LEFT JOIN pms_product p ON r.product_id = p.id
	// WHERE r.flash_promotion_id = @flashPromotionId AND r.flash_promotion_session_id = @flashPromotionSessionId
	// ORDER BY r.sort DESC) t)
	GetListCount(flashPromotionId, flashPromotionSessionId int64) (int64, error)
	// sql(SELECT r.id, r.flash_promotion_price, r.flash_promotion_count, r.flash_promotion_limit,
	//            r.flash_promotion_id, r.flash_promotion_session_id, r.product_id, r.sort,
	//            p.`name` name, p.product_sn product_sn, p.price price, p.stock stock
	// FROM sms_flash_promotion_product_relation r
	// LEFT JOIN pms_product p ON r.product_id = p.id
	// WHERE r.flash_promotion_id = @flashPromotionId AND r.flash_promotion_session_id = @flashPromotionSessionId
	// ORDER BY r.sort DESC OFFSET @offset LIMIT @limit)
	GetListData(flashPromotionId, flashPromotionSessionId int64, offset, limit int) ([]*dto.SmsFlashPromotionProductRelationDto, error)
}

func Generator() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./sms/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})
	db, _ := gorm.Open(mysql.Open("root:password@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)
	g.ApplyInterface(func() {}, model.SmsCouponHistory{})
	g.ApplyInterface(func() {}, model.SmsCoupon{})
	g.ApplyInterface(func() {}, model.SmsCouponProductRelation{})
	g.ApplyInterface(func() {}, model.SmsCouponProductCategoryRelation{})
	g.ApplyInterface(func(method SmsFlashPromotionProductRelationMethod) {}, model.SmsFlashPromotionProductRelation{})
	g.ApplyInterface(func() {}, model.SmsFlashPromotion{})
	g.ApplyInterface(func() {}, model.SmsFlashPromotionSession{})
	g.ApplyInterface(func() {}, model.SmsHomeAdvertise{})
	g.ApplyInterface(func() {}, model.SmsHomeBrand{})
	g.ApplyInterface(func() {}, model.SmsHomeNewProduct{})
	g.ApplyInterface(func() {}, model.SmsHomeRecommendProduct{})
	g.ApplyInterface(func() {}, model.SmsHomeRecommendSubject{})
	g.Execute()
}
