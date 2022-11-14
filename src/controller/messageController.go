package controller

import (
	"errors"

	repo "github.com/makks129/project-paper-planes/src/repository"
	"github.com/makks129/project-paper-planes/src/repository/db/model"
)

func GetMessage(userId string) (*model.Message, error) {
	assignedMessage, _ := repo.GetAssignedUnreadMessage(userId)
	if assignedMessage != nil {
		return assignedMessage, nil
	}
	latestMessage, _ := repo.GetLatestUnassignedMessage()
	if latestMessage != nil {
		repo.AssignMessage(userId)
		return latestMessage, nil
	}
	return nil, errors.New("no message available")
}
