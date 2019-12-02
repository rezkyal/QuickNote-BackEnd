package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/rezkyal/QuickNote-BackEnd/controllers"
	"github.com/rezkyal/QuickNote-BackEnd/middleware"
	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func main() {
	r := gin.Default()
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Panic(err)
	}

	r.Use(sessions.Sessions("mysession", store))

	var userController controllers.UserController
	var noteController controllers.NoteController
	userController.Init(db)
	noteController.Init(db)

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	r.GET("/api", userController.CreateUser)
	r.GET("/api/readAllNote/:username", noteController.ReadAllNote)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth())
	{
		authorized.POST("/api/readSearchNote", noteController.ReadSearchNote)
		authorized.GET("/api/createOneNote", noteController.CreateOneNote)
		authorized.POST("/api/readOneNote", noteController.ReadOneNote)
		authorized.POST("/api/updateOneNote", noteController.UpdateOneNote)
		authorized.POST("/api/deleteOneNote", noteController.DeleteOneNote)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	r.Run(":" + port)

}
