package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	MobileNumber string `gorm:"uniqueIndex;not null;size:20"`
}
