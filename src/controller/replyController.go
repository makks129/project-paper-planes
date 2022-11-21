package controller

import (
	"errors"

	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/model"
	repo "github.com/makks129/project-paper-planes/src/repository"
	"gorm.io/gorm"
)

func GetReplies(tx *gorm.DB, userId string) ([]*model.Reply, error) {
	replies, error := repo.GetUnreadReplies(tx, userId)

	// log.Println("GetReplies", "\n| replies: ", replies, "\n| ERROR: ", error, "\n ")

	if len(replies) > 0 {
		return replies, nil
	} else if !errors.As(error, &err.NotFoundError{}) {
		return nil, error
	}
	return nil, err.NothingAvailableError{}
}

func SaveReply(tx *gorm.DB, userId string, messageId uint, messageUserId string, text string) error {
	// Saving a reply for a message automatically acks the message as well
	ackMessageError := repo.AckMessage(tx, userId, messageId)
	if ackMessageError != nil {
		return ackMessageError
	}
	saveReplyError := repo.SaveReply(tx, userId, messageId, messageUserId, text)
	if saveReplyError != nil {
		return saveReplyError
	}
	return nil
}

func AckReply(userId string, replyId uint) error {
	return repo.AckReply(userId, replyId)
}
