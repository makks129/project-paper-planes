package router

import (
	"github.com/gin-gonic/gin"
)

const COOKIE_USER_ID = "user_id"

func SetupRouter(app *gin.Engine) {
	app.POST("/start", RequireCookie(COOKIE_USER_ID), PostStart)
	app.POST("/send-message", RequireCookie(COOKIE_USER_ID), ValidateBody[SendMessageBody], SendMessage)
	app.POST("/send-reply", RequireCookie(COOKIE_USER_ID), ValidateBody[SendReplyBody], SendReply)
	app.POST("/ack", RequireCookie(COOKIE_USER_ID), ValidateBody[AckBody], Ack)
}
