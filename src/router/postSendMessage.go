package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/err"
)

type SendMessageBody struct {
	Text string `json:"text" validate:"required,min=1,max=10000"`
}

type SendMessageErrorResponseBody struct {
	Code int `json:"code"`
}

func SendMessage(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	body := c.MustGet(VALIDATED_BODY).(*SendMessageBody)

	error := controller.SaveMessage(userId, body.Text)

	switch {
	case errors.As(error, &err.CannotWriteMoreMessagesError{}):
		c.JSON(http.StatusInternalServerError, gin.H{"code": err.CODE_CANNOT_WRITE_MORE_MESSAGES})
	case error != nil:
		// TODO log
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
	default:
		c.JSON(http.StatusCreated, gin.H{})
	}
}
