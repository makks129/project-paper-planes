package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/src/validator"
	"github.com/makks129/project-paper-planes/test/suit"
	testUtils "github.com/makks129/project-paper-planes/test/utils"
	"github.com/stretchr/testify/assert"
)

type TestRequestBody struct {
	Message string `json:"message" validate:"required,min=2,max=10"`
}

type ErrorResponseBody struct {
	Errors []validator.ValidationError `json:"errors"`
}

type TestResponseBody struct {
	ValidatedBody TestRequestBody `json:"validated_body"`
}

func Test_ValidateBody(t *testing.T) {
	app := InitApp()

	app.GET("/test", router.ValidateBody[TestRequestBody], func(c *gin.Context) {
		validatedBody, _ := c.Get(router.VALIDATED_BODY)
		c.JSON(http.StatusOK, gin.H{"validated_body": validatedBody})
	})

	s := suit.Of(&suit.SubTests{T: t})

	s.Test("returns 200 and validated body, if body is correct", func(t *testing.T) {
		w := httptest.NewRecorder()

		sendTestRequest(`{ "message": "Foo bar" }`, app, w)

		body := testUtils.FromJson[TestResponseBody](w.Body)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "Foo bar", body.ValidatedBody.Message)
	})

	s.Test("returns 400, if Content-Type is not application/json", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte("{}")))
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, w.Body.String(), `{"error":"Content-Type must be application/json"}`)
	})

	s.Test("returns 400 and validation errors, if body is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()

		sendTestRequest(`{ "invalid": "x" }`, app, w)

		body := testUtils.FromJson[ErrorResponseBody](w.Body)

		assert.Equal(t, 400, w.Code)
		assert.Len(t, body.Errors, 1)
		assert.Equal(t, "Message", body.Errors[0].Field)
		assert.Equal(t, "required", body.Errors[0].Tag)
	})

	s.Test("returns 400 and validation errors, if body has errors", func(t *testing.T) {
		w := httptest.NewRecorder()

		sendTestRequest(`{ "message": "1" }`, app, w)

		body := testUtils.FromJson[ErrorResponseBody](w.Body)

		assert.Equal(t, 400, w.Code)
		assert.Len(t, body.Errors, 1)
		assert.Equal(t, "Message", body.Errors[0].Field)
		assert.Equal(t, "min", body.Errors[0].Tag)
	})
}

func sendTestRequest(jsonStr string, app *gin.Engine, w *httptest.ResponseRecorder) {
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(w, req)
}
