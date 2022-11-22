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
	assignedUnreadMessage, err1 := repo.GetAssignedUnreadMessage(tx, userId)

	if assignedUnreadMessage != nil {
		return assignedUnreadMessage, nil
	} else if !errors.As(err1, &err.NotFoundError{}) {
		return nil, err1
	} // if such message is not found, continue

	assignedReadTodayMessage, err2 := repo.GetAssignedTodayMessage(tx, userId)

	if assignedReadTodayMessage != nil {
		return nil, err.CannotReceiveMoreMessagesError{}
	} else if !errors.As(err2, &err.NotFoundError{}) {
		return nil, err2
	} // if such message is not found, continue

	latestMessage, err3 := repo.GetLatestUnassignedMessage(tx, userId)

	if latestMessage != nil {

		err4 := repo.AssignMessage(userId, latestMessage.ID, tx)
		if err4 != nil {
			return nil, err4
		}

		return latestMessage, nil
	} else if !errors.As(err3, &err.NotFoundError{}) {
		return nil, err3
	}

	return nil, err.NothingAvailableError{}
}

func SaveMessage(userId string, text string) error {
	return repo.SaveMessage(userId, text)
}

func AckMessage(userId string, messageId uint) error {
	return repo.AckMessage(db.Db, userId, messageId)
}
