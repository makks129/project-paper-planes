package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/makks129/project-paper-planes/src/db"
// 	"github.com/makks129/project-paper-planes/src/model"
// 	"github.com/makks129/project-paper-planes/test/suit"
// 	"github.com/makks129/project-paper-planes/test/utils"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )

// func Test_PostStart_Cookie(t *testing.T) {
// 	app := InitApp()

// 	s := suit.Of(&suit.SubTests{T: t})

// 	s.Test("returns 400, if no user_id cookie", func(t *testing.T) {
// 		w := httptest.NewRecorder()
// 		req, _ := http.NewRequest("POST", "/start", nil)
// 		// no cookie
// 		app.ServeHTTP(w, req)

// 		assert.Equal(t, 400, w.Code)
// 	})
// }

// func Test_PostStart_NoContent(t *testing.T) {
// 	app := InitApp()
// 	db.InitDb()
// 	db.RunDbMigrations()

// 	s := suit.Of(&suit.SubTests{T: t})

// 	s.Test("returns 204, if no replies or messages exist", func(t *testing.T) {
// 		w := sendStartRequest(app)

// 		assert.Equal(t, 204, w.Code)
// 		assert.Equal(t, "", w.Body.String())
// 	})
// }

// func Test_PostStart_Replies(t *testing.T) {
// 	app := InitApp()
// 	db.InitDb()
// 	db.RunDbMigrations()

// 	cleanupDb := func() {
// 		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
// 		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
// 	}

// 	s := suit.Of(&suit.SubTests{
// 		T:          t,
// 		BeforeEach: cleanupDb,
// 		AfterAll:   cleanupDb,
// 	})

// 	s.Test("returns no replies, if no messages exist", func(t *testing.T) {
// 		w := sendStartRequest(app)

// 		assert.Equal(t, 204, w.Code)
// 		assert.Equal(t, "", w.Body.String())
// 	})

// 	s.Test("returns no replies, if replies do not exist", func(t *testing.T) {
// 		_ALICE_ID := ALICE_ID
// 		CreateMessage(BOB_ID, &_ALICE_ID, false)

// 		w := sendStartRequest(app)
// 		body := utils.FromJson[PostStartBody](w.Body)

// 		assert.Equal(t, 200, w.Code)
// 		assert.NotNil(t, body.Message)
// 	})

// 	s.Test("returns no replies, if replies exist but are already read", func(t *testing.T) {
// 		_BOB_ID := BOB_ID
// 		msg := CreateMessage(ALICE_ID, &_BOB_ID, false)

// 		reply := model.Reply{}
// 		db.Db.Create(&model.Reply{
// 			UserId:    BOB_ID,
// 			MessageId: msg.ID,
// 			Text:      "Reply to Lorem ipsum",
// 			IsRead:    true,
// 		}).First(&reply)

// 		w := sendStartRequest(app)

// 		assert.Equal(t, 204, w.Code)
// 		assert.Equal(t, "", w.Body.String())
// 	})

// 	s.Test("returns all available replies, if unread replies exist", func(t *testing.T) {
// 		_BOB_ID := BOB_ID
// 		msg1 := CreateMessage(ALICE_ID, &_BOB_ID, false)
// 		msg2 := CreateMessage(ALICE_ID, &_BOB_ID, false)
// 		msg3 := CreateMessage(ALICE_ID, &_BOB_ID, false)
// 		CreateReply(BOB_ID, msg1.ID, true)
// 		CreateReply(BOB_ID, msg2.ID, false)
// 		CreateReply(BOB_ID, msg3.ID, false)

// 		w := sendStartRequest(app)
// 		body := utils.FromJson[PostStartBody](w.Body)

// 		assert.Equal(t, 200, w.Code)
// 		assert.Len(t, body.Replies, 2)
// 		assert.Equal(t, msg2.ID, body.Replies[0].MessageId)
// 		assert.Equal(t, msg3.ID, body.Replies[1].MessageId)
// 	})
// }

// func Test_PostStart_Messages(t *testing.T) {
// 	app := initApp()
// 	db.InitDb()
// 	db.RunDbMigrations()

// 	cleanupDb := func() {
// 		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Message{})
// 		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reply{})
// 	}

// 	s := suit.Of(&suit.SubTests{
// 		T:          t,
// 		BeforeEach: cleanupDb,
// 		AfterAll:   cleanupDb,
// 	})

// 	s.Test("returns message, if assigned message exists", func(t *testing.T) {
// 		_ALICE_ID := ALICE_ID
// 		CreateMessage(BOB_ID, &_ALICE_ID, false)

// 		w := sendStartRequest(app)
// 		body := utils.FromJson[PostStartBody](w.Body)

// 		assert.Equal(t, 200, w.Code)
// 		assert.NotNil(t, body.Message)
// 	})

// 	s.Test("doesn't return message, if assigned message exists but it's already read", func(t *testing.T) {
// 		_ALICE_ID := ALICE_ID
// 		CreateMessage(BOB_ID, &_ALICE_ID, true)

// 		w := sendStartRequest(app)

// 		assert.Equal(t, 204, w.Code)
// 		assert.Equal(t, "", w.Body.String())
// 	})

// 	s.Test("returns message, if assigned-unread message doesn't exist and unassigned one exists", func(t *testing.T) {
// 		CreateMessage(BOB_ID, nil, false)

// 		w := sendStartRequest(app)
// 		body := utils.FromJson[PostStartBody](w.Body)

// 		assert.Equal(t, 200, w.Code)
// 		assert.NotNil(t, body.Message)
// 	})

// 	s.Test("doesn't return message, if neither assigned-unread or unassigned messages exist", func(t *testing.T) {
// 		_ALICE_ID := ALICE_ID
// 		CreateMessage(BOB_ID, &_ALICE_ID, true)

// 		w := sendStartRequest(app)

// 		assert.Equal(t, 204, w.Code)
// 		assert.NotNil(t, "", w.Body.String())
// 	})

// }

// type PostStartBody struct {
// 	Replies []*model.Reply `json:"replies"`
// 	Message *model.Message `json:"message"`
// }

// func sendSendMessageRequest(app *gin.Engine) *httptest.ResponseRecorder {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/send-message", nil)
// 	req.AddCookie(&http.Cookie{Name: "user_id", Value: ALICE_ID, Secure: true, HttpOnly: true})
// 	app.ServeHTTP(w, req)
// 	return w
// }