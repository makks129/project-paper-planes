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

const ALICE_ID = "mock_alice_id" // Alice sends messages
const BOB_ID = "mock_bob_id"     // Bob reads and replies to Alice's messages

func initApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	return app
}

func Test_GetStart_Cookie(t *testing.T) {
	app := initApp()

	s := suit.Of(&suit.SubTests{T: t})

	s.Test("returns 400, if no user_id cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/start", nil)
		// no cookie
		app.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func Test_GetStart_NoContent(t *testing.T) {
	app := initApp()
	db.InitDb()
	db.RunDbMigrations()

	s := suit.Of(&suit.SubTests{T: t})

	s.Test("returns 204, if no replies or messages exist", func(t *testing.T) {
		w := sendStartRequest(app)

		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})
}

func Test_GetStart_Replies(t *testing.T) {
	app := initApp()
	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns no replies, if no messages exist", func(t *testing.T) {
		w := sendStartRequest(app)

		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	s.Test("returns no replies, if replies do not exist", func(t *testing.T) {
		createMessage(BOB_ID, ALICE_ID, false)

		w := sendStartRequest(app)

		assert.Equal(t, 200, w.Code)

		var body GetStartBody
		json.Unmarshal(w.Body.Bytes(), &body)

		assert.NotNil(t, body.Message)
	})

	s.Test("returns no replies, if replies exist but are already read", func(t *testing.T) {
		msg := createMessage(ALICE_ID, BOB_ID, false)

		reply := model.Reply{}
		db.Db.Create(&model.Reply{
			UserId:    BOB_ID,
			MessageId: msg.ID,
			Text:      "Reply to Lorem ipsum",
			IsRead:    true,
		}).First(&reply)

		w := sendStartRequest(app)

		assert.Equal(t, 204, w.Code)
		assert.Equal(t, "", w.Body.String())
	})

	s.Test("returns all available replies, if unread replies exist", func(t *testing.T) {
		msg1 := createMessage(ALICE_ID, BOB_ID, false)
		msg2 := createMessage(ALICE_ID, BOB_ID, false)
		msg3 := createMessage(ALICE_ID, BOB_ID, false)
		createReply(BOB_ID, msg1.ID, true)
		createReply(BOB_ID, msg2.ID, false)
		createReply(BOB_ID, msg3.ID, false)

		w := sendStartRequest(app)

		assert.Equal(t, 200, w.Code)

		var body GetStartBody
		json.Unmarshal(w.Body.Bytes(), &body)

		assert.Len(t, body.Replies, 2)
		assert.Equal(t, msg2.ID, body.Replies[0].MessageId)
		assert.Equal(t, msg3.ID, body.Replies[1].MessageId)
	})
}

func Test_GetStart_Messages(t *testing.T) {
	// app := initApp()
	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns message, if assigned message exists", func(t *testing.T) {
		// w := sendStartRequest(app)

		// assert.Equal(t, 204, w.Code)

		assert.Equal(t, 1, 1)
	})

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

type GetStartBody struct {
	Replies []*model.Reply `json:"replies"`
	Message *model.Message `json:"message"`
}

func sendStartRequest(app *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/start", nil)
	req.AddCookie(&http.Cookie{Name: "user_id", Value: ALICE_ID, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}

func createMessage(userId string, assignedToUserId string, isRead bool) model.Message {
	msg := model.Message{}
	db.Db.Create(&model.Message{
		UserId:           userId,
		Text:             "Lorem ipsum",
		AssignedToUserId: assignedToUserId,
		AssignedAt:       time.Now(),
		IsRead:           isRead,
	}).First(&msg)
	return msg
}

func createReply(userId string, messageId uint, isRead bool) model.Reply {
	reply := model.Reply{}
	db.Db.Create(&model.Reply{
		UserId:    userId,
		MessageId: messageId,
		Text:      "Reply to Lorem ipsum",
		IsRead:    isRead,
	}).First(&reply)
	return reply
}
