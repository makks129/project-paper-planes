package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
