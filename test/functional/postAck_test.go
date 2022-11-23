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

func Test_PostAck(t *testing.T) {
	app := InitApp()
	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		deleteTables(model.Message{}, model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns 200, if message is acked", func(t *testing.T) {
		_BOB_ID := BOB_ID
		msg := CreateMessage(ALICE_ID, &_BOB_ID, false)

		json := fmt.Sprintf(`{"id": %d, "type": "message"}`, msg.ID)
		w := SendAckRequest(app, BOB_ID, json)

		assert.Equal(t, 200, w.Code)

		var message *model.Message
		res := db.Db.Table("messages").Take(&message)

		assert.Nil(t, res.Error)

		messageMatcher := mock.MatchedBy(func(m *model.Message) bool {
			return m.UserId == ALICE_ID &&
				m.AssignedToUserId.Valid &&
				m.AssignedToUserId.String == BOB_ID &&
				m.AssignedAt.Valid &&
				m.IsRead
		})

		assert.Equal(t, true, messageMatcher.Matches(message))
	})

	s.Test("returns 200, if reply is acked", func(t *testing.T) {
		_BOB_ID := BOB_ID
		msg := CreateMessage(ALICE_ID, &_BOB_ID, true)
		rpl := CreateReply(BOB_ID, msg.ID, ALICE_ID, false)

		json := fmt.Sprintf(`{"id": %d, "type": "reply"}`, rpl.ID)
		w := SendAckRequest(app, ALICE_ID, json)

		assert.Equal(t, 200, w.Code)

		var reply *model.Reply
		res := db.Db.Table("replies").Take(&reply)

		assert.Nil(t, res.Error)

		replyMatcher := mock.MatchedBy(func(m *model.Reply) bool {
			return m.UserId == BOB_ID &&
				m.AssignedToUserId.Valid &&
				m.AssignedToUserId.String == ALICE_ID &&
				m.AssignedAt.Valid &&
				m.IsRead
		})

		assert.Equal(t, true, replyMatcher.Matches(reply))
	})

	s.Test("returns 400, if type is invalid", func(t *testing.T) {
		w := SendAckRequest(app, ALICE_ID, `{"id": 1, "type": "invalid"}`)

		assert.Equal(t, 400, w.Code)
	})
}

func SendAckRequest(app *gin.Engine, fromUserId string, jsonStr string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/ack", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "user_id", Value: fromUserId, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
