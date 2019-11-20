package queryfunction

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rezkyal/QuickNote-BackEnd/models"
)

type NoteQuery struct {
	db *gorm.DB
}

func (n *NoteQuery) Init(db *gorm.DB) {
	n.db = db
}

func (n *NoteQuery) CreateNote(username string) models.Note {
	nt := models.Note{Username: username, CreatedOn: time.Now()}
	err := n.db.Create(&nt)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return nt
}

func (n *NoteQuery) FindNote(note_id int64) models.Note {
	var note models.Note
	err := n.db.Where("note_id", note_id).First(&note)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return note
}

func (n *NoteQuery) UpdateNote(note models.Note) models.Note {
	err := n.db.Update(&note)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return note
}

func (n *NoteQuery) DeleteNote(note models.Note) models.Note {
	err := n.db.Delete(&note)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return note
}
