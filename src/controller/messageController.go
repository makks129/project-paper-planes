package controller

import (
	"errors"

	"github.com/makks129/project-paper-planes/src/err"
	repo "github.com/makks129/project-paper-planes/src/repository"
	"github.com/makks129/project-paper-planes/src/repository/db/model"
)

func GetMessageOnStart(userId string) (*model.Message, error) {
	assignedMessage, err1 := repo.GetAssignedUnreadMessage(userId)

	if assignedMessage != nil {
		return assignedMessage, nil
	} else if !errors.As(err1, &err.NotFoundError{}) {
		return nil, err1
	}

	latestMessage, err2 := repo.GetLatestUnassignedMessage()

	if latestMessage != nil {

		err3 := repo.AssignMessage(userId, latestMessage.ID)
		if err3 != nil {
			return nil, err3
		}

		return latestMessage, nil
	} else if !errors.As(err2, &err.NotFoundError{}) {
		return nil, err2
	}

	return nil, err.NothingAvailableError{}
}
