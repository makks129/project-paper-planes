package controller

import (
	"errors"

	"github.com/makks129/project-paper-planes/src/db"
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	repo "github.com/makks129/project-paper-planes/src/repository"
	"gorm.io/gorm"
)

func GetMessageOnStart(tx *gorm.DB, userId string) (*model.Message, error) {
	assignedMessage, err1 := repo.GetAssignedUnreadMessage(tx, userId)

	if assignedMessage != nil {
		return assignedMessage, nil
	} else if !errors.As(err1, &err.NotFoundError{}) {
		return nil, err1
	}

	latestMessage, err2 := repo.GetLatestUnassignedMessage(tx, userId)

	if latestMessage != nil {

		err3 := repo.AssignMessage(userId, latestMessage.ID, tx)
		if err3 != nil {
			return nil, err3
		}

		return latestMessage, nil
	} else if !errors.As(err2, &err.NotFoundError{}) {
		return nil, err2
	}

	return nil, err.NothingAvailableError{}
}

func SaveMessage(userId string, text string) error {
	return repo.SaveMessage(userId, text)
}

func AckMessage(userId string, messageId uint) error {
	return repo.AckMessage(db.Db, userId, messageId)
}
