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
		bobbyDropTables(model.Message{}, model.Reply{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns 200, if reply is saved", func(t *testing.T) {
		_BOB_ID := BOB_ID
		msg := CreateMessage(ALICE_ID, &_BOB_ID, false)

		json := fmt.Sprintf(`{"message_id": %d, "message_user_id": "%s", "text": "Thank you for this message stranger."}`, msg.ID, msg.UserId)
		res := SendSendReplyRequest(app, BOB_ID, json)

		assert.Equal(t, 201, res.Code)

		var reply *model.Reply
		replyRes := db.Db.Table("replies").Take(&reply)

		assert.Nil(t, replyRes.Error)

		replyMatcher := mock.MatchedBy(func(r *model.Reply) bool {
			return r.UserId == BOB_ID &&
				r.MessageId == msg.ID &&
				r.Text == "Thank you for this message stranger." &&
				r.AssignedToUserId.Valid &&
				r.AssignedToUserId.String == ALICE_ID &&
				r.AssignedAt.Valid &&
				!r.IsRead
		})

		assert.Equal(t, true, replyMatcher.Matches(reply))

		var message *model.Message
		messageRes := db.Db.Table("messages").Take(&message)

		assert.Nil(t, messageRes.Error)

		messageMatcher := mock.MatchedBy(func(m *model.Message) bool {
			return m.UserId == ALICE_ID &&
				m.AssignedToUserId.Valid &&
				m.AssignedToUserId.String == BOB_ID &&
				m.AssignedAt.Valid &&
				m.IsRead // acked - marked as read
		})

		assert.Equal(t, true, messageMatcher.Matches(message))
	})

}

func SendSendReplyRequest(app *gin.Engine, fromUserId string, jsonStr string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/send-reply", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "user_id", Value: fromUserId, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
