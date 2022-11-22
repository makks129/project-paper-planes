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

const ALICE_ID = "alice_id"     // Alice: sends messages
const BOB_ID = "bob_id"         // Bob: send messages to Alice, or reads and replies to Alice's messages
const CHARLIE_ID = "charlie_id" // Charlie: send messages to Alice, or reads and replies to Alice's messages

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
	res := db.Db.Create(createMsg).First(&msg)
	if res.Error != nil {
		panic(res.Error)
	}
	return msg
}

func CreateReply(userId string, messageId uint, assignedToUserId string, isRead bool) model.Reply {
	createReply := &model.Reply{
		UserId:           userId,
		MessageId:        messageId,
		Text:             "Reply to Lorem ipsum",
		AssignedToUserId: sql.NullString{String: assignedToUserId, Valid: true},
		AssignedAt:       sql.NullTime{Time: time.Now(), Valid: true},
		IsRead:           isRead,
	}

	reply := model.Reply{}
	res := db.Db.Create(createReply).First(&reply)
	if res.Error != nil {
		panic(res.Error)
	}
	return reply
}
