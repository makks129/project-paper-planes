package controller

import (
	"errors"

	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	repo "github.com/makks129/project-paper-planes/src/repository"
	"gorm.io/gorm"
)

func GetReplies(userId string, tx *gorm.DB) ([]*model.Reply, error) {
	replies, error := repo.GetUnreadReplies(userId, tx)

	// log.Println("GetReplies", "\n| replies: ", replies, "\n| ERROR: ", error, "\n ")

	if len(replies) > 0 {
		return replies, nil
	} else if !errors.As(error, &err.NotFoundError{}) {
		return nil, error
	}
	return nil, err.NothingAvailableError{}
}

func SaveReply(userId string, messageId uint, text string) error {
	return repo.SaveReply(userId, messageId, text)
}
