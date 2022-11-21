package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/err"
)

type AckBody struct {
	Id   uint   `json:"id" validate:"required"`
	Type string `json:"type" validate:"eq=message|eq=reply"`
}

func Ack(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	body := c.MustGet(VALIDATED_BODY).(*AckBody)

	var error error
	switch body.Type {
	case "message":
		error = controller.AckMessage(userId, body.Id)
	case "reply":
		error = controller.AckReply(userId, body.Id)
	}

	log.Println("ack", "\n| ERROR: ", error, "\n ")

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
