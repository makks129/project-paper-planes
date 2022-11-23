package main

import (
	"testing"
	"time"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/makks129/project-paper-planes/test/utils"
	"github.com/stretchr/testify/assert"
)

func Test_WriteOnlyOneMessagePerDay(t *testing.T) {
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

	s.Test("user doesn't get to write another message, if they have already written a message today", func(t *testing.T) {
		// Alice sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, ALICE_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Alice sends a message (2)
		msgJson2 := `{"text": "I hope you will have a great day."}`
		sendMessageRes2 := SendSendMessageRequest(app, ALICE_ID, msgJson2)
		sendMessageRes2Body := utils.FromJson[router.SendMessageErrorResponseBody](sendMessageRes2.Body)

		assert.Equal(t, 500, sendMessageRes2.Code)
		assert.Equal(t, 30, sendMessageRes2Body.Code)
	})

	s.Test("user gets to write another message, if it's already the next day", func(t *testing.T) {
		// Alice sends a message
		msgJson := `{"text": "Hey there stranger."}`
		sendMessageRes := SendSendMessageRequest(app, ALICE_ID, msgJson)

		assert.Equal(t, 201, sendMessageRes.Code)

		// Find and modify Alice's message
		yesterday := time.Now().UTC().AddDate(0, 0, -1)
		updateRes := db.Db.Table("messages").
			Where("user_id = ?", ALICE_ID).
			Updates(model.Message{
				CreatedAt: yesterday,
			})
		assert.Nil(t, updateRes.Error)

		// Alice sends a message (2)
		msgJson2 := `{"text": "I hope you will have a great day."}`
		sendMessageRes2 := SendSendMessageRequest(app, ALICE_ID, msgJson2)

		assert.Equal(t, 201, sendMessageRes2.Code)
	})

}
