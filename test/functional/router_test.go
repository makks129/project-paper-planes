package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/repository/db"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/stretchr/testify/assert"
)

func TestGetStart(t *testing.T) {

	app := gin.Default()
	router.SetupRouter(app)

	db.InitDb()

	t.Run("returns reply, if reply exists", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		// assert.Equal(t, "pong", w.Body.String())

	})

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
