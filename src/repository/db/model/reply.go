package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	MessageId uint `gorm:"not null"`
	//...
}
