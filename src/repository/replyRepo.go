package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	"gorm.io/gorm"
)

func GetUnreadReplies(userId string, tx *gorm.DB) ([]*model.Reply, error) {
	replies := []*model.Reply{}
	res := tx.Table("replies AS r").
		Select("r.*, m.text AS message_text, m.created_at AS message_created_at").
		Joins("LEFT JOIN messages AS m ON r.message_id = m.id").
		Where("r.assigned_to_user_id = ? AND r.is_read = ?", userId, false).
		Find(&replies)

	log.Println("GetUnreadReplies", "\n| replies: ", replies, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case len(replies) == 0:
		return nil, err.NotFoundError{}
	case res.Error == nil:
		return replies, nil
	default:
		return nil, res.Error
	}
}

func SaveReply(userId string, messageId uint, messageUserId string, text string) error {
	res := db.Db.Create(&model.Reply{
		UserId:           userId,
		MessageId:        messageId,
		Text:             text,
		AssignedToUserId: sql.NullString{String: messageUserId, Valid: true},
		AssignedAt:       sql.NullTime{Time: time.Now(), Valid: true},
		IsRead:           false,
	})

	log.Println("SaveReply", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func AckReply(userId string, replyId uint) error {
	updates := model.Reply{IsRead: true}
	res := db.Db.Table("replies").Where("id = ? AND assigned_to_user_id = ?", replyId, userId).Updates(updates)

	log.Println("AckReply", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}
