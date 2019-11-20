package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rezkyal/QuickNote-BackEnd/queryfunction"

	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func main() {
	router := gin.Default()
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var userQuery queryfunction.UserQuery
	var noteQuery queryfunction.NoteQuery
	userQuery.Init(db)
	noteQuery.Init(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)

	userQuery.FindOrCreateUser("hiyahiya")
}
