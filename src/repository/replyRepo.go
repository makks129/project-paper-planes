package repositories

import (
	"database/sql"
	"time"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	"github.com/makks129/project-paper-planes/src/utils"
	"gorm.io/gorm"
)

func GetUnreadReplies(tx *gorm.DB, userId string) ([]*model.Reply, error) {
	replies := []*model.Reply{}
	res := tx.Table("replies AS r").
		Select("r.*, m.text AS message_text, m.created_at AS message_created_at").
		Joins("LEFT JOIN messages AS m ON r.message_id = m.id").
		Where("r.assigned_to_user_id = ? AND r.is_read = ?", userId, false).
		Find(&replies)

	utils.Log("GetUnreadReplies", "\n| replies: ", replies, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case len(replies) == 0:
		return nil, err.NotFoundError{}
	case res.Error == nil:
		return replies, nil
	default:
		return nil, res.Error
	}
}

func SaveReply(tx *gorm.DB, userId string, messageId uint, messageUserId string, text string) error {
	res := tx.Create(&model.Reply{
		UserId:           userId,
		MessageId:        messageId,
		Text:             text,
		AssignedToUserId: sql.NullString{String: messageUserId, Valid: true},
		AssignedAt:       sql.NullTime{Time: time.Now().UTC(), Valid: true},
		IsRead:           false,
	})

	utils.Log("SaveReply", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func AckReply(userId string, replyId uint) error {
	updates := model.Reply{IsRead: true}
	res := db.Db.Table("replies").Where("id = ? AND assigned_to_user_id = ?", replyId, userId).Updates(updates)

	utils.Log("AckReply", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}
