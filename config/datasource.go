package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type DataSource struct {
	Url      string
	Username string
	Password string
}

func InitDataSource() {
	var err error
	var dataSource DataSource
	err = settings.UnmarshalKey("datasource", &dataSource)
	if err != nil {
		log.Fatalf("get config failed: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@%s", dataSource.Username, dataSource.Password, dataSource.Url)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}
}

func GetDb() *gorm.DB {
	return db
}
