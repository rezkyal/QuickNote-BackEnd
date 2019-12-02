package models

import "time"

type User struct {
	Username   string `gorm:"PRIMARY_KEY;UNIQUE"`
	Password   string `gorm:"DEFAULT NULL"`
	CreatedOn  time.Time
	NotesOwned []Note `gorm:"foreignkey:Username;association_foreignkey:Username"`
}
