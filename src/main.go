package main

import (
	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/src/task"
)

func main() {
	db.InitDb()
	db.RunDbMigrations()

	go task.RunRecurringMessageUnassignTask()

	app := gin.Default()
	router.SetupRouter(app)
	app.Run(":9000")
}
