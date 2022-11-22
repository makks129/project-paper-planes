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

func Test_StandardConversation(t *testing.T) {
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

	s.Test("performs standard conversation", func(t *testing.T) {
		// Alice starts (1)
		startAliceRes := SendStartRequest(app, ALICE_ID)
		startAliceResBody := utils.FromJson[router.PostStartResponseBody](startAliceRes.Body)
		// gets nothing
		assert.Equal(t, 200, startAliceRes.Code)
		assert.Equal(t, 10, startAliceResBody.Code)

		// Alice sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, ALICE_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Bob starts (1)
		startBobRes := SendStartRequest(app, BOB_ID)
		startBobResBody := utils.FromJson[router.PostStartResponseBody](startBobRes.Body)
		// gets Alice's message
		assert.Equal(t, 200, startBobRes.Code)
		assert.Equal(t, "Hey there stranger.", startBobResBody.Message.Text)

		// Bob replies to Alice's message
		rplJson := fmt.Sprintf(`{`+
			`"message_id": %d, `+
			`"message_user_id": "%s", `+
			`"text": "Hey stranger, I hope you have a great day."}`,
			startBobResBody.Message.ID,
			startBobResBody.Message.UserId)
		sendReplyRes := SendSendReplyRequest(app, BOB_ID, rplJson)

		assert.Equal(t, 201, sendReplyRes.Code)

		// Alice starts (2)
		startAliceRes2 := SendStartRequest(app, ALICE_ID)
		startAliceRes2Body := utils.FromJson[router.PostStartResponseBody](startAliceRes2.Body)
		// gets Bob's reply
		assert.Equal(t, 200, startAliceRes2.Code)
		assert.Len(t, startAliceRes2Body.Replies, 1)
		assert.Equal(t, "Hey stranger, I hope you have a great day.", startAliceRes2Body.Replies[0].Text)

		// Alice acks Bob's reply
		ackJson := fmt.Sprintf(`{"id": %d, "type": "reply"}`, startAliceRes2Body.Replies[0].ID)
		ackRes := SendAckRequest(app, ALICE_ID, ackJson)

		assert.Equal(t, 200, ackRes.Code)

		// Alice starts (3)
		startAliceRes3 := SendStartRequest(app, ALICE_ID)
		startAliceRes3Body := utils.FromJson[router.PostStartResponseBody](startAliceRes3.Body)
		// gets nothing
		assert.Equal(t, 200, startAliceRes3.Code)
		assert.Equal(t, 10, startAliceRes3Body.Code)

		// Bob starts (2)
		startBobRes2 := SendStartRequest(app, BOB_ID)
		startBobRes2Body := utils.FromJson[router.PostStartResponseBody](startBobRes2.Body)
		// gets nothing
		assert.Equal(t, 200, startBobRes2.Code)
		assert.Equal(t, 20, startBobRes2Body.Code)
	})

}
