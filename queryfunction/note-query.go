package queryfunction

import (
	"github.com/jinzhu/gorm"
)

type NoteQuery struct {
	db *gorm.DB
}

func (n *NoteQuery) init(db *gorm.DB) {
	n.db = db
}
