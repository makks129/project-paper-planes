package main

import (
	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/router"
)

func main() {
	db.InitDb()
	db.RunDbMigrations()

	app := gin.Default()

	router.SetupRouter(app)

	app.Run(":9000")

	print("ololo test")

	// TODO start message free-up loop
}
