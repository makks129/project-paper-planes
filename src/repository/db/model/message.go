package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserId           string    `gorm:"type:varchar(36);not null"`
	Text             string    `gorm:"type:text;not null"`
	AssignedToUserId string    `gorm:"type:varchar(36)"`
	AssignedAt       time.Time `gorm:"type:datetime"`
	IsRead           bool      `gorm:"type:boolean;not null;default:false"`
}
