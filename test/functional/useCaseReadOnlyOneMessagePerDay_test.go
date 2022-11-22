package main

import (
	"fmt"
	"testing"

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

	s.Test("user doesn't get to read another message, if they already read a message today", func(t *testing.T) {
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

	// TODO if user doesn't ack then they should keep receiving the same message during the whole day
	// TODO the next day they should receive a new message

}
