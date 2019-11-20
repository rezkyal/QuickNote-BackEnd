package models

import "time"

type User struct {
	UserID     int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT; UNIQUE"`
	Username   string `gorm: "UNIQUE"`
	Password   string `gorm: "DEFAULT NULL"`
	CreatedOn  time.Time
	NotesOwned []Note `gorm:"foreignkey:UserID"`
}
