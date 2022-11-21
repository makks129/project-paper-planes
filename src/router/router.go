package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/utils"
	"gorm.io/gorm"
)

const COOKIE_USER_ID = "user_id"

func SetupRouter(app *gin.Engine) {
	app.POST("/start", RequireCookie(COOKIE_USER_ID), postStart)
	app.POST("/send-message", RequireCookie(COOKIE_USER_ID), ValidateBody[SendMessageBody], sendMessage)
	app.POST("/send-reply", RequireCookie(COOKIE_USER_ID), ValidateBody[SendReplyBody], sendReply)
	app.POST("/ack-message", ackMessage)
}

func postStart(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	error := db.Db.Transaction(func(tx *gorm.DB) error {

		replies, err1 := controller.GetReplies(userId, tx)
		// log.Println("GetStart", "\n| replies: ", replies, "\n| ERROR: ", err1, "\n ")
		switch {
		case len(replies) > 0:
			c.JSON(http.StatusOK, gin.H{"replies": replies})
			return nil
		case errors.As(err1, &err.NothingAvailableError{}):
			break
		default:
			utils.Error(err1)
			return err.GenericServerError{}
		}

		message, err2 := controller.GetMessageOnStart(userId, tx)
		// log.Println("GetStart", "\n| message: ", message, "\n| ERROR: ", err2, "\n ")
		switch {
		case message != nil:
			c.JSON(http.StatusOK, gin.H{"message": message})
			return nil
		case errors.As(err2, &err.NothingAvailableError{}):
			c.JSON(http.StatusNoContent, gin.H{})
			return nil
		default:
			return err.GenericServerError{}
		}
	})

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}
}

type SendMessageBody struct {
	Text string `json:"text" validate:"required,min=1,max=10000"`
}

func sendMessage(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	body := c.MustGet(VALIDATED_BODY).(*SendMessageBody)

	error := controller.SaveMessage(userId, body.Text)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type SendReplyBody struct {
	MessageId uint   `json:"message_id" validate:"required"`
	Text      string `json:"text" validate:"required,min=1,max=10000"`
}

func sendReply(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	body := c.MustGet(VALIDATED_BODY).(*SendReplyBody)

	error := controller.SaveReply(userId, body.MessageId, body.Text)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func ackMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
