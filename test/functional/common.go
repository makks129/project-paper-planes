package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/router"
	"gorm.io/gorm"
)

const ALICE_ID = "mock_alice_id" // Alice: sends messages
const BOB_ID = "mock_bob_id"     // Bob: send messages to Alice, or reads and replies to Alice's messages

func InitApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	return app
}

// https://xkcd.com/327/
func bobbyDropTables(models ...interface{}) {
	for _, m := range models {
		db.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(m)
	}
}

func CreateMessage(userId string, assignedToUserId *string, isRead bool) model.Message {
	createMsg := &model.Message{
		UserId:     userId,
		Text:       "Lorem ipsum",
		AssignedAt: sql.NullTime{Time: time.Now(), Valid: true},
		IsRead:     isRead,
	}
	if assignedToUserId != nil {
		createMsg.AssignedToUserId = sql.NullString{String: *assignedToUserId, Valid: true}
	} else {
		createMsg.AssignedToUserId = sql.NullString{Valid: false}
	}

	msg := model.Message{}
	db.Db.Create(createMsg).First(&msg)
	return msg
}

func CreateReply(userId string, messageId uint, assignedToUserId *string, isRead bool) model.Reply {
	createReply := &model.Reply{
		UserId:     userId,
		MessageId:  messageId,
		Text:       "Reply to Lorem ipsum",
		AssignedAt: sql.NullTime{Time: time.Now(), Valid: true},
		IsRead:     isRead,
	}
	if assignedToUserId != nil {
		createReply.AssignedToUserId = sql.NullString{String: *assignedToUserId, Valid: true}
	} else {
		createReply.AssignedToUserId = sql.NullString{Valid: false}
	}

	reply := model.Reply{}
	db.Db.Create(createReply).First(&reply)
	return reply
}
