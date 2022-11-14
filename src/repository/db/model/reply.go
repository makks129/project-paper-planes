package model

import (
	"time"
)

type Reply struct {
	ID uint `gorm:"column:id;primarykey;type:bigint unsigned;not null"`

	MessageId string `gorm:"column:message_id;type:bigint unsigned;not null"`
	Text      string `gorm:"column:text;type:text;not null"`

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
