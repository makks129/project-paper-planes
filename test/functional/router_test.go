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

const ALICE_ID = "mock_alice_id"
const BOB_ID = "mock_bob_id"

func TestGetStart(t *testing.T) {

	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Recovery())
	router.SetupRouter(app)

	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		// AfterAll:   cleanupDb,
	})

	// Cookie

	s.Skip("returns 400, if no user_id cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/start", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	// No content

	s.Skip("returns 204, if no replies or messages exist", func(t *testing.T) {
		w := sendStartRequest(app)
		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	// Replies

	s.Skip("returns no replies, if no messages exist", func(t *testing.T) {
		w := sendStartRequest(app)
		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	s.Skip("returns no replies, if replies do not exist", func(t *testing.T) {
		db.Db.Create(&model.Message{
			UserId:           BOB_ID,
			Text:             "Lorem ipsum",
			AssignedToUserId: ALICE_ID,
			AssignedAt:       time.Now(),
			IsRead:           false,
		})

		w := sendStartRequest(app)

		assert.Equal(t, 200, w.Code)

		var msgBody MessageBody
		json.Unmarshal(w.Body.Bytes(), &msgBody)

		assert.NotNil(t, msgBody.Message)
	})

	s.Test("returns no replies, if replies exist but are already read", func(t *testing.T) {
		msg := model.Message{}
		db.Db.Create(&model.Message{
			UserId:           ALICE_ID,
			Text:             "Lorem ipsum",
			AssignedToUserId: BOB_ID,
			AssignedAt:       time.Now(),
			IsRead:           false,
		}).First(&msg)

		reply := model.Reply{}
		db.Db.Create(&model.Reply{
			UserId:    BOB_ID,
			MessageId: msg.ID,
			Text:      "Reply to Lorem ipsum",
			IsRead:    false,
		}).First(&reply)

		// TODO
	})

	// s.Test("returns all available replies, if unread replies exist", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// // Messages

	// s.Test("returns message, if assigned message exists", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.Test("doesn't return message, if assigned message exists but it's already read", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.Test("returns message, if assigned-unread message doesn't exist and unassigned one exists", func(t *testing.T) {
	// 	// TODO
	// 	assert.Equal(t, 1, 1)
	// })

	// s.Test("doesn't return message, if neither assigned-unread or unassigned messages exist", func(t *testing.T) {
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
	req.AddCookie(&http.Cookie{Name: "user_id", Value: ALICE_ID, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
