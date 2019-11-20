package models

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	NoteID    int64 `gorm:"PRIMARY_KEY;UNIQUE;AUTO_INCREMENT"`
	UserID    int64
	Title     string
	Note      string
	CreatedOn string
}
