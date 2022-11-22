package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/utils"
	"gorm.io/gorm"
)

const NO_CONTENT_CODE_NOTHING_AVAILABLE = 10
const NO_CONTENT_CODE_CANNOT_RECEIVE_MORE_MESSAGES = 20

type PostStartResponseBody struct {
	Replies []*model.Reply `json:"replies"`
	Message *model.Message `json:"message"`
	Code    int            `json:"code"`
}

func PostStart(c *gin.Context) {
	userIdCookie, _ := c.Request.Cookie(COOKIE_USER_ID)
	userId := userIdCookie.Value

	error := db.Db.Transaction(func(tx *gorm.DB) error {

		replies, err1 := controller.GetReplies(tx, userId)
		// log.Println("postStart", "\n| replies: ", replies, "\n| ERROR: ", err1, "\n ")
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

		message, err2 := controller.GetMessageOnStart(tx, userId)
		// log.Println("postStart", "\n| message: ", message, "\n| ERROR: ", err2, "\n ")
		switch {
		case message != nil:
			c.JSON(http.StatusOK, gin.H{"message": message})
			return nil
		case errors.As(err2, &err.NothingAvailableError{}):
			c.JSON(http.StatusOK, gin.H{"code": NO_CONTENT_CODE_NOTHING_AVAILABLE})
			return nil
		case errors.As(err2, &err.CannotReceiveMoreMessagesError{}):
			c.JSON(http.StatusOK, gin.H{"code": NO_CONTENT_CODE_CANNOT_RECEIVE_MORE_MESSAGES})
			return nil
		default:
			return err.GenericServerError{}
		}
	})

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}
}
