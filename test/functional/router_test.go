package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/stretchr/testify/assert"
)

const MOCK_USER_ID = "mock_user_id"

func TestGetStart(t *testing.T) {

	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	db.InitDb()
	db.RunDbMigrations()

	// Cookie

	t.Run("returns 400, if no user_id cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/start", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	// No content

	t.Run("returns 204, if no replies or messages exist", func(t *testing.T) {
		w := sendStartRequest(app)
		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	// Replies

	// t.Run("returns no replies, if no messages exist", func(t *testing.T) {
	// 	// db.Db.Create(&model.Reply{
	// 	// 	MessageId: "1",
	// 	// 	Text:      "Lorem ipsum",
	// 	// })

	// 	w := httptest.NewRecorder()
	// 	req, _ := http.NewRequest("GET", "/start", nil)
	// 	app.ServeHTTP(w, req)

	// 	assert.Equal(t, 200, w.Code)
	// 	// assert.Equal(t, "pong", w.Body.String())
	// })

	t.Run("returns no replies, if replies do not exist", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("returns no replies, if replies exist but are already read", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("returns all available replies, if unread replies exist", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	// Messages

	t.Run("returns message, if assigned message exists", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("doesn't return message, if assigned message exists but it's already read", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("returns message, if assigned-unread message doesn't exist and unassigned one exists", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

	t.Run("doesn't return message, if neither assigned-unread or unassigned messages exist", func(t *testing.T) {
		// TODO
		assert.Equal(t, 1, 1)
	})

}

func sendStartRequest(app *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/start", nil)
	req.AddCookie(&http.Cookie{Name: "user_id", Value: MOCK_USER_ID, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
