package model

import (
	"time"
)

// TODO indices

type Message struct {
	ID uint `gorm:"column:id;primarykey"`

	UserId           string    `gorm:"column:user_id;type:varchar(36);not null"`
	Text             string    `gorm:"column:text;type:text;not null"`
	AssignedToUserId string    `gorm:"column:assigned_to_user_id;type:varchar(36)"`
	AssignedAt       time.Time `gorm:"column:assigned_at;type:datetime"`
	IsRead           bool      `gorm:"column:is_read;type:boolean;not null;default:false"`

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
