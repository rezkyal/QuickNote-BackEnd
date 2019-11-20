package models

import "time"

type Note struct {
	NoteID    int64 `gorm:"PRIMARY_KEY;UNIQUE;AUTO_INCREMENT"`
	Username  string
	Title     string
	Note      string
	CreatedOn time.Time
	UpdatedOn time.Time
}
