package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/utils"
)

func SetupRouter(app *gin.Engine) {
	app.GET("/start", getStart)
	app.POST("/send-message", sendMessage)
	app.POST("/send-reply", sendReply)
	app.POST("/ack-message", ackMessage)
}

func getStart(c *gin.Context) {
	userId := utils.GetCookie("user_id", c.Request)

	reply, _ := controller.GetReply(userId)

	if reply != nil {
		c.JSON(http.StatusOK, gin.H{"reply": reply})
		return
	}

	message, _ := controller.GetMessage(userId)

	if message != nil {
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "nothing available"})
}

func sendMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func sendReply(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ackMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
