package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Datasource struct {
	Url      string
	Username string
	Password string
}

func init() {
	var err error
	var datasource Datasource
	_ = settings.UnmarshalKey("datasource", &datasource)
	dsn := fmt.Sprintf("%s:%s@%s", datasource.Username, datasource.Password, datasource.Url)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}
}

func GetDb() *gorm.DB {
	return db
}
