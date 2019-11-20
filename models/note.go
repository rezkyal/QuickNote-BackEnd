package models

type Note struct {
	NoteID    int64  `json:"note_id"`
	UserID    int64  `json:"user_id"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	CreatedOn string `json:"created_on"`
}
