package main

import (
	"log"

	"github.com/rezkyal/QuickNote-BackEnd/util"
)

func main() {
	// router := gin.Default()
	db, err := util.GetDB()

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	router.GET("/noteDetail")
	router.GET("/")
}
