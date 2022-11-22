package main

import (
	"testing"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/router"
	"github.com/makks129/project-paper-planes/test/suit"
	"github.com/makks129/project-paper-planes/test/utils"
	"github.com/stretchr/testify/assert"
)

func Test_DisallowGettingOwnMessage(t *testing.T) {
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

	s.Test("user doesn't get their own message", func(t *testing.T) {
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

		// Alice starts (2)
		startAliceRes2 := SendStartRequest(app, ALICE_ID)
		startAliceRes2Body := utils.FromJson[router.PostStartResponseBody](startAliceRes2.Body)
		// gets nothing
		assert.Equal(t, 200, startAliceRes2.Code)
		assert.Equal(t, 10, startAliceRes2Body.Code)
	})

}
