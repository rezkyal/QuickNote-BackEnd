package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/rezkyal/QuickNote-BackEnd/controllers"
	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func main() {
	router := gin.Default()
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var userController controllers.UserController
	var noteController controllers.NoteController
	userController.Init(db)
	noteController.Init(db)

	router.GET("/readAllNote/:username", noteController.ReadAllNote)
	router.POST("/readSearchNote", noteController.ReadSearchNote)
	router.GET("/createOneNote/:username", noteController.CreateOneNote)
	router.POST("/readOneNote", noteController.ReadOneNote)
	router.POST("/updateOneNote", noteController.UpdateOneNote)
	router.POST("/deleteOneNote", noteController.DeleteOneNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	router.Run(":" + port)

}
