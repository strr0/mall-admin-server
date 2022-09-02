package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mall-admin-server/ums/query"
)

func init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}
	query.SetDefault(db)
}
