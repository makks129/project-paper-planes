package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/utils"
	"github.com/makks129/project-paper-planes/src/validator"
)

const VALIDATED_BODY = "validated_body"

func ValidateBody[T interface{}](c *gin.Context) {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
		c.Abort()
		return
	}

	validatedBody, validationErrors, error := validator.ValidateRequestBody[T](c)

	switch {
	case error != nil:
		utils.Log(error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.GenericServerError{}.Error()})
		c.Abort()
	case len(validationErrors) > 0:
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		c.Abort()
	default:
		c.Set(VALIDATED_BODY, validatedBody)
	}
}
