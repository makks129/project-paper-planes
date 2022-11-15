package controller

import (
	"github.com/makks129/project-paper-planes/src/err"
	"github.com/makks129/project-paper-planes/src/repository/db/model"
	"gorm.io/gorm"
)

func GetReply(userId string, tx *gorm.DB) (*model.Reply, error) {
	return nil, err.NothingAvailableError{}
}
