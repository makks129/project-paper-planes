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

func GetAssignedUnreadMessage(tx *gorm.DB, userId string) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").
		Where("assigned_to_user_id = ?", userId). // assigned
		Where("is_read = ?", false).              // unread
		Take(&message)

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

func GetAssignedTodayMessage(tx *gorm.DB, userId string) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").
		Where("assigned_to_user_id = ?", userId). // assigned
		Where("is_read = ?", true).               // read
		Where("DATE(assigned_at) = DATE(NOW())"). // today
		Take(&message)

	log.Println("GetAssignedTodayMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

	switch {
	case res.Error == nil:
		return message, nil
	case res.Error.Error() == "record not found":
		return nil, err.NotFoundError{}
	default:
		return nil, res.Error
	}
}

func GetLatestUnassignedMessage(tx *gorm.DB, userId string) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").
		Where("user_id != ?", userId).        // from someone else
		Where("assigned_to_user_id IS NULL"). // unassigned
		Order("created_at DESC").             // latest
		Take(&message)

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
	updates := model.Message{
		AssignedToUserId: sql.NullString{String: userId, Valid: true},
		AssignedAt:       sql.NullTime{Time: time.Now(), Valid: true},
	}
	res := tx.Table("messages").Where("id = ?", messageId).Updates(updates)

	log.Println("AssignMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func SaveMessage(userId string, text string) error {
	res := db.Db.Create(&model.Message{
		UserId: userId,
		Text:   text,
		IsRead: false,
	})

	log.Println("SaveMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func AckMessage(tx *gorm.DB, userId string, messageId uint) error {
	updates := model.Message{IsRead: true}
	res := tx.Table("messages").Where("id = ? AND assigned_to_user_id = ?", messageId, userId).Updates(updates)

	log.Println("AckMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}
