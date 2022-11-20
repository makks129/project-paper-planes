package validator

import (
	"sync"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var once sync.Once

func Validator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func ValidateRequestBody[T interface{}](c *gin.Context) (*T, []ValidationError, error) {
	var body T
	err1 := c.ShouldBindJSON(&body)
	if err1 != nil {
		return nil, nil, err1
	}

	err2 := Validator().Struct(body)

	res := []ValidationError{}
	if err2 != nil {
		validationErrors := err2.(validator.ValidationErrors)
		for _, e := range validationErrors {
			res = append(res, ValidationError{Field: e.Field(), Tag: e.Tag()})
		}
	}

	if len(res) == 0 {
		return &body, nil, nil
	}

	return nil, res, nil
}
