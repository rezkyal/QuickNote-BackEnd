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
	nt := models.Note{Username: username, CreatedOn: time.Now().UTC(), UpdatedOn: time.Now().UTC()}
	err := n.db.Create(&nt)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return nt
}

func (n *NoteQuery) FindNote(note_id int64) models.Note {
	var note models.Note
	err := n.db.Where("note_id = ?", note_id).First(&note)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return note
}

func (n *NoteQuery) FindNoteByQuery(username string, query string) []models.Note {
	var notes []models.Note
	state := n.db.Where("Username = ? AND (LOWER(Title) LIKE ? or LOWER(Note) LIKE ?)", username, "%"+query+"%", "%"+query+"%").Order("created_on asc").Find(&notes)
	if state.Error != nil {
		if gorm.IsRecordNotFoundError(state.Error) {
			return notes
		} else {
			log.Panic(state.Error)
		}
	}
	return notes
}

func (n *NoteQuery) UpdateNote(note models.Note) models.Note {
	note.UpdatedOn = time.Now().UTC()
	err := n.db.Save(&note)
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
