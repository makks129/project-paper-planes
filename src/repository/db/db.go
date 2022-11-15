package db

import (
	"fmt"
	"os"

	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"github.com/makks129/project-paper-planes/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	goEnv := os.Getenv("GO_ENV")

	// TODO enable logger only in dev
	// dbLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,
	// 		LogLevel:                  logger.Info,
	// 		IgnoreRecordNotFoundError: false,
	// 		Colorful:                  true,
	// 	},
	// )

	utils.Log("InitDb --")

	var dsn string
	switch goEnv {
	case "test":
		dsn = "root:root@tcp(0.0.0.0:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
	default:
		dsn = "ppp.user:ppp123@tcp(mysql:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
	}

	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger:                 dbLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("Failed to connect to DB")
		panic(err)
	}
	println("Connected to MySQL DB")
	Db = mysqlDb

	utils.Log("InitDb connected to mysql db")

	//
	// test
	//
	RunDbMigrations()
	res := &[]string{}
	Db.Raw("show tables").Scan(&res)
	fmt.Println(">>> RES", res)
}

func RunDbMigrations() {
	Db.AutoMigrate(&model.Message{})
	Db.AutoMigrate(&model.Reply{})
}
