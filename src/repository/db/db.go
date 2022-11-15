package db

import (
	"log"
	"os"
	"time"

	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"github.com/makks129/project-paper-planes/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() {
	var dsn string
	var dbLogger logger.Interface

	switch os.Getenv("GO_ENV") {
	case "test":
		dsn = "root:root@tcp(0.0.0.0:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
		dbLogger = logger.Default.LogMode(logger.Silent)
	default:
		dsn = "ppp.user:ppp123@tcp(mysql:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
		dbLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		)
	}

	config := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	if dbLogger != nil {
		config.Logger = dbLogger
	}

	mysqlDb, err := gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		utils.Log("Failed to connect to DB")
		panic(err)
	}
	utils.Log("Connected to MySQL DB")
	Db = mysqlDb
}

func RunDbMigrations() {
	Db.AutoMigrate(&model.Message{})
	Db.AutoMigrate(&model.Reply{})
}
