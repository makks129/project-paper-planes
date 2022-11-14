package db

import (
	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	dsn := "ppp.user:ppp123@tcp(mysql:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	println("Connected to MySQL DB")
	Db = mysqlDb
}

func RunDbMigrations() {
	Db.AutoMigrate(&model.Message{})
	Db.AutoMigrate(&model.Reply{})
}
