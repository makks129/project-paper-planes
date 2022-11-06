package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/makks129/project-paper-planes/src/lib"
)

var db *gorm.DB

func setupDb() {
	dsn := "ppp.user:ppp123@tcp(mysql:3306)/ppp?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Connected to db")
	db = mysqlDb
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	setupDb()

	r := setupRouter()

	lib.Foo()

	r.Run(":9000")
}
