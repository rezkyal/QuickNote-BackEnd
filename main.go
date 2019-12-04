package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/rezkyal/QuickNote-BackEnd/controllers"
	"github.com/rezkyal/QuickNote-BackEnd/middleware"
	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("qu1ckn0t3"))
	if err != nil {
		log.Panic(err)
	}

	r.Use(sessions.Sessions("mysession", store))

	var userController controllers.UserController
	var noteController controllers.NoteController
	userController.Init(db)
	noteController.Init(db)

	r.GET("/api/user", userController.CreateUser)
	r.GET("/api/user/initUser/:username", userController.InitUser)
	r.POST("/api/user/login", userController.Login)
	r.GET("/api/user/logout", userController.Logout)
	r.POST("/api/user/setNewPassword", userController.SetNewPassword)

	authorized := r.Group("/api")
	authorized.Use(middleware.Auth())
	{
		authorized.POST("/user/changePassword", userController.ChangePassword)

		note := authorized.Group("/note")
		note.GET("/readAllNote", noteController.ReadAllNote)
		note.POST("/readSearchNote", noteController.ReadSearchNote)
		note.GET("/createOneNote", noteController.CreateOneNote)
		note.POST("/readOneNote", noteController.ReadOneNote)
		note.POST("/updateOneNote", noteController.UpdateOneNote)
		note.POST("/deleteOneNote", noteController.DeleteOneNote)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	r.Run(":" + port)

}
