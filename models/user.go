package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID    int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT; UNIQUE"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
}
