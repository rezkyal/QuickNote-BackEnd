package queryfunction

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/rezkyal/QuickNote-BackEnd/models"

	"github.com/jinzhu/gorm"

	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func initQuery() (NoteQuery, *gorm.DB) {
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}

	var noteQuery NoteQuery
	noteQuery.Init(db)

	return noteQuery, db
}

func TesCreate(t *testing.T) {
	noteQuery, db := initQuery()
	noteQuery.CreateNote("admin")
	db.Close()
}

func TestRead(t *testing.T) {
	NoteQuery, db := initQuery()
	note := NoteQuery.FindNote(1)
	if note.Title != "this is a testing note" {
		t.Errorf("Read data testing failed")
	}
	db.Close()
}

func TestUpdate(t *testing.T) {
	NoteQuery, db := initQuery()
	note := NoteQuery.FindNote(1)
	note.Title = "this is an update test"
	note.Note = "this is an update test"
	note.UpdatedOn = time.Date(2019, 11, 24, 10, 50, 0, 0, time.UTC)
	fmt.Printf("title %s, note %s\n", note.Title, note.Note)
	fmt.Println(note.UpdatedOn)
	NoteQuery.UpdateNote(note)
	note = models.Note{}

	note = NoteQuery.FindNote(1)
	if note.Title != "this is an update test" {
		t.Errorf("Update data first test failed")
	}
	if note.Note != "this is an update test" {
		t.Errorf("Update data second test failed")
	}
	if util.CustomFormat(note.UpdatedOn) != util.CustomFormat(time.Date(2019, 11, 24, 10, 50, 0, 0, time.UTC)) {
		t.Errorf("Update data third test failed")
	}
	note.Title = "this is a testing note"
	note.Note = "this is a testing note"
	note.UpdatedOn = time.Date(2019, 11, 21, 10, 35, 58, 949124, time.UTC)
	NoteQuery.UpdateNote(note)
	note = models.Note{}

	note = NoteQuery.FindNote(1)
	if note.Title != "this is a testing note" {
		t.Errorf("Update data fourth test failed")
	}
	if note.Note != "this is a testing note" {
		t.Errorf("Update data fifth test failed")
	}
	if util.CustomFormat(note.UpdatedOn) != util.CustomFormat(time.Date(2019, 11, 21, 10, 35, 58, 949124, time.UTC)) {
		t.Errorf("Update data sixth test failed")
	}
	db.Close()
}
