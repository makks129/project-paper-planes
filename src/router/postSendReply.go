package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"gorm.io/gorm"
)

type SendReplyBody struct {
	MessageId     uint   `json:"message_id" validate:"required"`
	MessageUserId string `json:"message_user_id" validate:"required"`
	Text          string `json:"text" validate:"required,min=1,max=10000"`
}

func SendReply(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	body := c.MustGet(VALIDATED_BODY).(*SendReplyBody)

	error := db.Db.Transaction(func(tx *gorm.DB) error {

		error := controller.SaveReply(tx, userId, body.MessageId, body.MessageUserId, body.Text)

		if error != nil {
			log.Println("SendReply", "\n| ERROR: ", error, "\n ")
			return err.GenericServerError{}
		}

		c.JSON(http.StatusOK, gin.H{})
		return nil
	})

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}
}
