package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO cover 500 case with test (mock gorm to throw error)

// TODO ack message on reply

func Test_PostSendReply(t *testing.T) {
	app := InitApp()
	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		bobbyDropTables(model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns 200, if reply is saved", func(t *testing.T) {
		json := fmt.Sprintf(`{"message_id": 42, "message_user_id": "%s", "text": "Answer to the Ultimate Question of Life, the Universe, and Everything"}`, BOB_ID)
		w := sendSendReplyRequest(app, json)

		assert.Equal(t, 200, w.Code)

		var reply *model.Reply
		res := db.Db.Table("replies").Take(&reply)

		assert.Nil(t, res.Error)

		replyMatcher := mock.MatchedBy(func(m *model.Reply) bool {
			return m.UserId == BOB_ID &&
				m.MessageId == 42 &&
				m.Text == "Answer to the Ultimate Question of Life, the Universe, and Everything" &&
				!m.IsRead
		})

		assert.Equal(t, true, replyMatcher.Matches(reply))
	})

}

func sendSendReplyRequest(app *gin.Engine, jsonStr string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/send-reply", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "user_id", Value: BOB_ID, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
