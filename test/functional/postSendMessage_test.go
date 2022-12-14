package main

import (
	"bytes"
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

func Test_PostSendMessage(t *testing.T) {
	app := InitApp()
	db.InitDb()
	db.RunDbMigrations()

	cleanupDb := func() {
		deleteTables(model.Message{})
	}

	s := suit.Of(&suit.SubTests{
		T:          t,
		BeforeEach: cleanupDb,
		AfterAll:   cleanupDb,
	})

	s.Test("returns 200, if message is saved", func(t *testing.T) {
		w := SendSendMessageRequest(app, ALICE_ID, `{"text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."}`)

		assert.Equal(t, 201, w.Code)

		var message *model.Message
		res := db.Db.Table("messages").Take(&message)

		assert.Nil(t, res.Error)

		messageMatcher := mock.MatchedBy(func(m *model.Message) bool {
			return m.UserId == ALICE_ID &&
				m.Text == "Lorem ipsum dolor sit amet, consectetur adipiscing elit." &&
				!m.AssignedToUserId.Valid &&
				!m.AssignedAt.Valid &&
				!m.IsRead
		})

		assert.Equal(t, true, messageMatcher.Matches(message))
	})

}

func SendSendMessageRequest(app *gin.Engine, fromUserId string, jsonStr string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("POST", "/send-message", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "user_id", Value: fromUserId, Secure: true, HttpOnly: true})
	app.ServeHTTP(w, req)
	return w
}
