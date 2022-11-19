package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/utils"
)

func RequireCookie(cookie string) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, error := c.Request.Cookie(cookie)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			c.Abort()
		}
	}
}

func ValidateBody[T interface{}](c *gin.Context) {
	validationErrors, error := utils.ValidateRequestBody[T](c)

	switch {
	case error != nil:
		utils.Log(error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
		c.Abort()
	case len(validationErrors) > 0:
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		c.Abort()
	}
}