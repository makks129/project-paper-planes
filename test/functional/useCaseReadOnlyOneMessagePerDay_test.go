package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/makks129/project-paper-planes/test/utils"
	"github.com/stretchr/testify/assert"
)

func Test_ReadOnlyOneMessagePerDay(t *testing.T) {
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

	s.Test("user doesn't get to read another message, if they have already read a message today", func(t *testing.T) {
		// Bob sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, BOB_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Alice starts (1)
		startAliceRes := SendStartRequest(app, ALICE_ID)
		startAliceResBody := utils.FromJson[router.PostStartResponseBody](startAliceRes.Body)
		// gets Bob's message
		assert.Equal(t, 200, startAliceRes.Code)
		assert.Equal(t, "Hey there stranger.", startAliceResBody.Message.Text)

		// Alice acks Bob's message
		ackJson := fmt.Sprintf(`{"id": %d, "type": "message"}`, startAliceResBody.Message.ID)
		ackRes := SendAckRequest(app, ALICE_ID, ackJson)

		assert.Equal(t, 200, ackRes.Code)

		// Charlie sends a message
		msgJson2 := `{"text": "I hope you have a great day."}`
		sendMessageRes2 := SendSendMessageRequest(app, CHARLIE_ID, msgJson2)

		assert.Equal(t, 201, sendMessageRes2.Code)

		// Alice starts (2)
		startAliceRes2 := SendStartRequest(app, ALICE_ID)
		startAliceRes2Body := utils.FromJson[router.PostStartResponseBody](startAliceRes2.Body)
		// doesn't get Charlie's message, because she has already read Bob's message today
		assert.Equal(t, 200, startAliceRes2.Code)
		assert.Equal(t, 20, startAliceRes2Body.Code)
	})

	s.Test("user will keep receiving the same message, if it's assigned to them but they don't ack it", func(t *testing.T) {
		// Bob sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, BOB_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Alice starts (1)
		startAliceRes := SendStartRequest(app, ALICE_ID)
		startAliceResBody := utils.FromJson[router.PostStartResponseBody](startAliceRes.Body)
		// gets Bob's message
		assert.Equal(t, 200, startAliceRes.Code)
		assert.Equal(t, "Hey there stranger.", startAliceResBody.Message.Text)

		// Alice starts (2)
		startAliceRes2 := SendStartRequest(app, ALICE_ID)
		startAliceResBody2 := utils.FromJson[router.PostStartResponseBody](startAliceRes2.Body)
		// gets Bob's message
		assert.Equal(t, 200, startAliceRes2.Code)
		assert.Equal(t, "Hey there stranger.", startAliceResBody2.Message.Text)
	})

	s.Test("user gets to read another message, if it's already the next day", func(t *testing.T) {
		// Bob sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, BOB_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Alice starts (1)
		startAliceRes := SendStartRequest(app, ALICE_ID)
		startAliceResBody := utils.FromJson[router.PostStartResponseBody](startAliceRes.Body)
		// gets Bob's message
		assert.Equal(t, 200, startAliceRes.Code)
		assert.Equal(t, "Hey there stranger.", startAliceResBody.Message.Text)

		// Alice acks Bob's message
		ackJson := fmt.Sprintf(`{"id": %d, "type": "message"}`, startAliceResBody.Message.ID)
		ackRes := SendAckRequest(app, ALICE_ID, ackJson)

		assert.Equal(t, 200, ackRes.Code)

		// Charlie sends a message
		msgJson2 := `{"text": "I hope you have a great day."}`
		sendMessageRes2 := SendSendMessageRequest(app, CHARLIE_ID, msgJson2)

		assert.Equal(t, 201, sendMessageRes2.Code)

		// Find and modify Bob's message
		yesterday := time.Now().AddDate(0, 0, -1)
		updateRes := db.Db.Table("messages").
			Where("user_id = ?", BOB_ID).
			Updates(model.Message{
				AssignedAt: sql.NullTime{Time: yesterday, Valid: true},
			})
		assert.Nil(t, updateRes.Error)

		// Alice starts (2)
		startAliceRes2 := SendStartRequest(app, ALICE_ID)
		startAliceRes2Body := utils.FromJson[router.PostStartResponseBody](startAliceRes2.Body)
		// gets Charlie's message
		assert.Equal(t, 200, startAliceRes2.Code)
		assert.Equal(t, "I hope you have a great day.", startAliceRes2Body.Message.Text)
	})

}
