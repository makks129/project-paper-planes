package repositories

import (
	"log"
	"time"

	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"gorm.io/gorm"
)

func GetAssignedUnreadMessage(userId string, tx *gorm.DB) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").Where("assigned_to_user_id = ? AND is_read = ?", userId, false).Take(&message)

	log.Println("GetAssignedUnreadMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case res.Error == nil:
		return message, nil
	case res.Error.Error() == "record not found":
		return nil, err.NotFoundError{}
	default:
		return nil, res.Error
	}
}

func GetLatestUnassignedMessage(tx *gorm.DB) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").Where("assigned_to_user_id IS NULL").Order("created_at DESC").Take(&message)

	log.Println("GetLatestUnassignedMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case res.Error == nil:
		return message, nil
	case res.Error.Error() == "record not found":
		return nil, err.NotFoundError{}
	default:
		return nil, res.Error
	}
}

func AssignMessage(userId string, messageId uint, tx *gorm.DB) error {
	updates := model.Message{AssignedToUserId: userId, AssignedAt: time.Now()}
	res := tx.Table("messages").Where("id = ?", messageId).Updates(updates)

	log.Println("AssignMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}
