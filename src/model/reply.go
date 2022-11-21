package model

import (
	"database/sql"
	"time"
)

// TODO indices

type Reply struct {
	ID uint `gorm:"column:id;primarykey;type:bigint unsigned;not null"`

	UserId           string         `gorm:"column:user_id;type:varchar(36);not null"`
	MessageId        uint           `gorm:"column:message_id;type:bigint unsigned;not null"`
	Text             string         `gorm:"column:text;type:text;not null"`
	AssignedToUserId sql.NullString `gorm:"column:assigned_to_user_id;type:varchar(36)"`
	AssignedAt       sql.NullTime   `gorm:"column:assigned_at;type:datetime"`
	IsRead           bool           `gorm:"column:is_read;type:boolean;not null;default:false"`

	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
