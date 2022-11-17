package repositories

import (
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

	// log.Println("GetUnreadReplies", "\n| replies: ", replies, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case len(replies) == 0:
		return nil, err.NotFoundError{}
	case res.Error == nil:
		return replies, nil
	default:
		return nil, res.Error
	}
}
