package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"gorm.io/gorm"
)

func SetupRouter(app *gin.Engine) {
	app.GET("/test", test)

	app.GET("/start", getStart)
	app.POST("/send-message", sendMessage)
	app.POST("/send-reply", sendReply)
	app.POST("/ack-message", ackMessage)
}

// TODO delete
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getStart(c *gin.Context) {
	userIdCookie, cookieError := c.Request.Cookie("user_id")
	if cookieError != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	userId := userIdCookie.Value

	error := db.Db.Transaction(func(tx *gorm.DB) error {

		reply, err1 := controller.GetReply(userId, tx)
		switch {
		case reply != nil:
			c.JSON(http.StatusOK, gin.H{"reply": reply})
			return nil
		case errors.As(err1, &err.NothingAvailableError{}):
			break
		default:
			return err1
		}

		message, err2 := controller.GetMessageOnStart(userId, tx)
		switch {
		case message != nil:
			c.JSON(http.StatusOK, gin.H{"message": message})
			return nil
		case errors.As(err2, &err.NothingAvailableError{}):
			c.JSON(http.StatusNoContent, gin.H{})
			return nil
		default:
			return err2
		}
	})

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}
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
