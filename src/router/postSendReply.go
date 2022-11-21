package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/err"
)

type SendReplyBody struct {
	MessageId uint   `json:"message_id" validate:"required"`
	Text      string `json:"text" validate:"required,min=1,max=10000"`
}

func SendReply(c *gin.Context) {
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
