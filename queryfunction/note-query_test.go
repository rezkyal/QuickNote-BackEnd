package queryfunction

import (
	"log"
	"testing"
	"time"

	"github.com/rezkyal/QuickNote-BackEnd/models"

	"github.com/jinzhu/gorm"

	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func initQuery() (NoteQuery, UserQuery, *gorm.DB) {
	db, err := util.GetDBTest()

	if err != nil {
		log.Panic(err)
	}

	var noteQuery NoteQuery
	var userQuery UserQuery
	noteQuery.Init(db)
	userQuery.Init(db)

	return noteQuery, userQuery, db
}

func TestCreateAndDelete(t *testing.T) {
	noteQuery, _, db := initQuery()
	defer db.Close()
	note := noteQuery.CreateNote("admin")
	noteID := note.NoteID

	note = models.Note{}
	note = noteQuery.FindNote(noteID)
	noteQuery.DeleteNote(note)

}

func TestRead(t *testing.T) {
	noteQuery, _, db := initQuery()
	defer db.Close()
	note := noteQuery.FindNote(1)
	if note.Title != "this is a testing note" {
		t.Errorf("Read data testing failed")
	}
}

func searchOne(t *testing.T, noteQuery NoteQuery, query string, expectedID []int64) {
	var notes []models.Note
	var noteID []int64

	notes = noteQuery.FindNoteByQuery("admin", query)
	for _, note := range notes {
		noteID = append(noteID, note.NoteID)
	}

	if len(noteID) != len(expectedID) {
		t.Errorf("Search data query: \"%s\" failed!, data length not equal!. Expected %d, got %d", query, len(expectedID), len(noteID))
		return
	}

	for _, i := range expectedID {
		check := util.CheckIntInArray(i, noteID)
		if !check {
			t.Errorf("NoteID %d not founded!", i)
		}
	}
}

func TestSearch(t *testing.T) {
	noteQuery, _, db := initQuery()
	defer db.Close()

	expectedID := []int64{1, 2, 3}
	query := "test"

	searchOne(t, noteQuery, query, expectedID)

	expectedID = []int64{2, 3}
	query = "look"

	searchOne(t, noteQuery, query, expectedID)

	expectedID = []int64{}
	query = "asasfasf"

	searchOne(t, noteQuery, query, expectedID)

}

func TestUpdate(t *testing.T) {
	noteQuery, _, db := initQuery()
	note := noteQuery.FindNote(1)
	note.Title = "this is an update test"
	note.Note = "this is an update test"
	note.UpdatedOn = time.Date(2019, 11, 24, 10, 50, 0, 0, time.UTC)
	noteQuery.UpdateNote(note)
	note = models.Note{}

	note = noteQuery.FindNote(1)
	if note.Title != "this is an update test" {
		t.Errorf("Update data first test failed (Title not updated)")
	}
	if note.Note != "this is an update test" {
		t.Errorf("Update data second test failed (Note not updated)")
	}
	if util.CustomFormat(note.UpdatedOn) != util.CustomFormat(time.Date(2019, 11, 24, 10, 50, 0, 0, time.UTC)) {
		t.Errorf("Update data third test failed (UpdatedOn not updated)")
	}
	note.Title = "this is a testing note"
	note.Note = "this is a testing note"
	note.UpdatedOn = time.Date(2019, 11, 21, 10, 35, 58, 949124, time.UTC)
	noteQuery.UpdateNote(note)
	note = models.Note{}

	note = noteQuery.FindNote(1)
	if note.Title != "this is a testing note" {
		t.Errorf("Update data fourth test failed (Title not updated)")
	}
	if note.Note != "this is a testing note" {
		t.Errorf("Update data fifth test failed (Note not updated)")
	}
	if util.CustomFormat(note.UpdatedOn) != util.CustomFormat(time.Date(2019, 11, 21, 10, 35, 58, 949124, time.UTC)) {
		t.Errorf("Update data sixth test failed (UpdatedOn not updated)")
	}
	db.Close()
}
