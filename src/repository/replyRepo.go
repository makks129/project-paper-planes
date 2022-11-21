package repositories

import (
	"log"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	"gorm.io/gorm"
)

func GetUnreadReplies(userId string, tx *gorm.DB) ([]*model.Reply, error) {
	replies := []*model.Reply{}

	// SELECT r.*
	// FROM messages m
	// LEFT JOIN replies r
	// ON m.id = r.message_id
	// WHERE m.user_id = ? AND r.is_read = 0
	res := tx.Table("messages").Select("replies.*").
		Joins("LEFT JOIN replies ON messages.id = replies.message_id").
		Where("messages.user_id = ?", userId).
		Where("replies.is_read = 0").
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

func SaveReply(userId string, messageId uint, text string) error {
	res := db.Db.Create(&model.Reply{
		UserId:    userId,
		MessageId: messageId,
		Text:      text,
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
