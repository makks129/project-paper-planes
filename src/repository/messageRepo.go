package repositories

import (
	"github.com/makks129/project-paper-planes/src/repository/db"
	"github.com/makks129/project-paper-planes/src/repository/db/model"
)

func GetAssignedUnreadMessage(userId string) (*model.Message, error) {
	var message model.Message
	res := db.Db.Table("messages").Where("user_id = ? AND is_read = ?", userId, false).First(&message)
	if res.Error != nil {
		return nil, res.Error
	}
	return &message, nil
}

func GetLatestUnassignedMessage() (*model.Message, error) {
	// TODO
	return nil, nil
}

func AssignMessage(userId string) {
	// TODO
}
