package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const MOCK_USER_ID = "mock_user_id"

func TestGetStart(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Recovery())
	router.SetupRouter(app)

	db.InitDb()
	db.RunDbMigrations()

	s := suit.Of(&suit.SubTests{T: t,
		BeforeEach: func(t *testing.T) {
			db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
			db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
		},
		AfterAll: func() {
			db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
			db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
		},
	})

	// Cookie

	s.TestIt("returns 400, if no user_id cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/start", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	s.TestIt("returns 400, if no user_id cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/start", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	// No content

	s.TestIt("returns 204, if no replies or messages exist", func(t *testing.T) {
		w := sendStartRequest(app)
		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	// Replies

	s.TestIt("returns no replies, if no messages exist", func(t *testing.T) {
		w := sendStartRequest(app)
		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	s.TestIt("returns no replies, if replies do not exist", func(t *testing.T) {
		db.Db.Create(&model.Message{
			UserId:           "some_other_user_id",
			Text:             "Lorem ipsum",
			AssignedToUserId: MOCK_USER_ID,
			AssignedAt:       time.Now(),
			IsRead:           false,
		})

		w := sendStartRequest(app)

		assert.Equal(t, 200, w.Code)

		var msgBody MessageBody
		json.Unmarshal(w.Body.Bytes(), &msgBody)

		assert.NotNil(t, msgBody.Message)
	})

	s.TestIt("returns no replies, if replies exist but are already read", func(t *testing.T) {
		// db.Db.Create(&model.Message{
		// 	UserId:           MOCK_USER_ID,
		// 	Text:             "Lorem ipsum",
		// 	AssignedToUserId: "some_other_user_id",
		// 	AssignedAt:       time.Now(),
		// 	IsRead:           false,
		// })
	})

	// s.TestIt("returns all available replies, if unread replies exist", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// // Messages

	// s.TestIt("returns message, if assigned message exists", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.TestIt("doesn't return message, if assigned message exists but it's already read", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.TestIt("returns message, if assigned-unread message doesn't exist and unassigned one exists", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.TestIt("doesn't return message, if neither assigned-unread or unassigned messages exist", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

}

type MessageBody struct {
	Message *model.Message `json:"message"`
}

func sendStartRequest(app *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/start", nil)
	req.AddCookie(&http.Cookie{Name: "user_id", Value: MOCK_USER_ID, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
