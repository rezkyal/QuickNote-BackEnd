package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rezkyal/QuickNote-BackEnd/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rezkyal/QuickNote-BackEnd/queryfunction"
	"github.com/rezkyal/QuickNote-BackEnd/util"
)

type NoteController struct {
	userQuery *queryfunction.UserQuery
	noteQuery *queryfunction.NoteQuery
}

func (n *NoteController) Init(db *gorm.DB) {
	n.userQuery = &queryfunction.UserQuery{}
	n.noteQuery = &queryfunction.NoteQuery{}

	n.userQuery.Init(db)
	n.noteQuery.Init(db)
}

func readNote(n *NoteController, c *gin.Context) (models.Note, error) {
	noteid := c.DefaultPostForm("noteid", "0")
	if noteid == "0" {
		return models.Note{}, fmt.Errorf("noteID field empty")
	}

	noteidint, err := strconv.ParseInt(noteid, 10, 64)
	if err != nil {
		panic(err)
	}

	note := n.noteQuery.FindNote(noteidint)
	return note, nil
}

func (n *NoteController) CreateOneNote(c *gin.Context) {
	username := c.Param("username")
	note := n.noteQuery.CreateNote(username)
	c.JSON(200, note)
}

func (n *NoteController) ReadAllNote(c *gin.Context) {
	username := c.Param("username")
	user := n.userQuery.FindOrCreateUser(username)
	for i := range user.NotesOwned {
		user.NotesOwned[i].Title = util.Ellipsis(user.NotesOwned[i].Title, 150)
		user.NotesOwned[i].Note = util.Ellipsis(user.NotesOwned[i].Note, 150)
	}

	c.JSON(200, user.NotesOwned)
}

func (n *NoteController) ReadOneNote(c *gin.Context) {
	note, err := readNote(n, c)
	if err != nil {
		c.JSON(400, err)
		log.Panic(err)
		return
	}
	c.JSON(200, note)
}

func (n *NoteController) UpdateOneNote(c *gin.Context) {
	note, err := readNote(n, c)
	if err != nil {
		c.JSON(400, err)
		log.Panic(err)
		return
	}

	title := c.DefaultPostForm("title", note.Title)
	notebody := c.DefaultPostForm("note", note.Note)
	note.Title = title
	note.Note = notebody

	n.noteQuery.UpdateNote(note)

	c.JSON(200, note)

}

func (n *NoteController) DeleteOneNote(c *gin.Context) {
	note, err := readNote(n, c)
	if err != nil {
		c.JSON(400, err)
		log.Panic(err)
		return
	}

	n.noteQuery.DeleteNote(note)
	c.JSON(200, gin.H{
		"success": "true"})
}
