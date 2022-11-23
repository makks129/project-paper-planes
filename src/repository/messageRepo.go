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

func GetAssignedUnreadMessage(tx *gorm.DB, userId string) (*model.Message, error) {
	var message *model.Message
	res := tx.Table("messages").
		Where("assigned_to_user_id = ?", userId). // assigned
		Where("is_read = ?", false).              // unread
		Take(&message)

	utils.Log("GetAssignedUnreadMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

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

	utils.Log("GetAssignedTodayMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

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

	utils.Log("GetLatestUnassignedMessage", "\n| message: ", message, "\n| ERROR: ", res.Error, "\n ")

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
		AssignedAt:       sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	res := tx.Table("messages").
		Where("id = ?", messageId).
		Updates(updates)

	utils.Log("AssignMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func SaveMessage(userId string, text string) error {
	res := db.Db.Create(&model.Message{
		UserId: userId,
		Text:   text,
		IsRead: false,
	})

	utils.Log("SaveMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func HasUserCreatedMessageToday(userId string) (bool, error) {
	var count int64
	res := db.Db.Table("messages").
		Where("user_id = ?", userId).            // from this user
		Where("DATE(created_at) = DATE(NOW())"). // created today
		Count(&count)

	utils.Log("HasUserCreatedMessageToday", "\n| count: ", count, "\n| ERROR: ", res.Error, "\n ")

	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}

func AckMessage(tx *gorm.DB, userId string, messageId uint) error {
	updates := model.Message{IsRead: true}
	res := tx.Table("messages").
		Where("id = ? AND assigned_to_user_id = ?", messageId, userId).
		Updates(updates)

	utils.Log("AckMessage", "\n| ERROR: ", res.Error, "\n ")

	return res.Error
}

func UnassignOldAssignedUnreadMessage() (int64, error) {
	res := db.Db.Table("messages").
		Where("assigned_at < DATE_SUB(NOW(), INTERVAL 1 DAY)"). // assigned more than 1 day ago
		Where("is_read = ?", false).                            // unread
		Updates(map[string]interface{}{
			"assigned_to_user_id": nil,
			"assigned_at":         nil,
		})

	utils.Log("UnassignAssignedUnreadMessage", "\n| rows: ", res.RowsAffected, "\n| ERROR: ", res.Error, "\n ")

	return res.RowsAffected, res.Error
}
