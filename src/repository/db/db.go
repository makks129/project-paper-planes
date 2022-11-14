package db

import (
	"log"
	"os"
	"time"

	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() {
	// TODO enable logger only in dev
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	dsn := "ppp.user:ppp123@tcp(mysql:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
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
