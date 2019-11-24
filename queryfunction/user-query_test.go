package queryfunction

import (
	"testing"

	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func TestFindUserAndNotesOwned(t *testing.T) {
	_, userQuery, db := initQuery()
	defer db.Close()

	var noteID []int64

	user := userQuery.FindOrCreateUser("admin")
	if user.Username != "admin" {
		t.Errorf("User admin not founded/created!")
	}

	for _, note := range user.NotesOwned {
		noteID = append(noteID, note.NoteID)
	}

	expectedID := []int64{1, 2, 3}

	if len(user.NotesOwned) != len(expectedID) {
		t.Errorf("Note owning test failed!, data length not equal!. Expected %d, got %d", len(expectedID), len(noteID))
		return
	}

	for _, i := range expectedID {
		check := util.CheckIntInArray(i, noteID)
		if !check {
			t.Errorf("NoteID %d not founded!", i)
		}
	}
}
