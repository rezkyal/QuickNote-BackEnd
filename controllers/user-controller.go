package controllers

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rezkyal/QuickNote-BackEnd/models"
	"github.com/rezkyal/QuickNote-BackEnd/queryfunction"
	"github.com/rezkyal/QuickNote-BackEnd/util"
)

type UserController struct {
	userQuery *queryfunction.UserQuery
}

func (u *UserController) Init(db *gorm.DB) {
	u.userQuery = &queryfunction.UserQuery{}
	u.userQuery.Init(db)
}

func (u *UserController) InitUser(c *gin.Context) {
	username := c.Param("username")
	user := u.userQuery.FindOrCreateUser(username)

	session := sessions.Default(c)

	checkname := ""

	if session.Get("username") != nil {
		checkname = session.Get("username").(string)
	}

	if checkname != username {
		session.Set("username", username)
		if user.Password == "" {
			session.Set("loggedin", true)
		} else {
			session.Set("loggedin", false)
		}
	}
	err := session.Save()
	if err != nil {
		log.Panic(err)
	}
	c.JSON(200, gin.H{"status": "1", "username": session.Get("username").(string)})
}

func (u *UserController) CreateUser(c *gin.Context) {
	username := util.RandomString(10)
	user, available := u.userQuery.CreateUser(username)
	if available {
		c.JSON(200, user)
	} else {
		u.CreateUser(c)
	}
}

func (u *UserController) SetNewPassword(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)

	user := u.userQuery.FindOrCreateUser(username)
	password := []byte(c.PostForm("password"))
	setPassword(u, user, password)
	c.JSON(200, gin.H{"success": "1"})
}

func (u *UserController) ChangePassword(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)

	user := u.userQuery.FindOrCreateUser(username)
	oldpassword := []byte(c.PostForm("oldpassword"))
	newpassword := []byte(c.PostForm("newpassword"))
	stats := checkPassword(u, user, oldpassword)

	if !stats {
		c.JSON(400, gin.H{"status": "0", "message": "Wrong old password!"})
	} else {
		setPassword(u, user, newpassword)
		c.JSON(200, gin.H{"status": "1"})
	}

}

func (u *UserController) Login(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)

	user := u.userQuery.FindOrCreateUser(username)
	password := []byte(c.PostForm("password"))
	stats := checkPassword(u, user, password)

	if !stats {
		c.JSON(400, gin.H{"status": "0", "message": "Wrong password!"})
	} else {
		session.Set("loggedin", true)

		err := session.Save()
		if err != nil {
			log.Panic(err)
		}

		c.JSON(200, gin.H{"status": "1"})
	}
}

func (u *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("loggedin", false)

	err := session.Save()
	if err != nil {
		log.Panic(err)
	}

	c.JSON(200, gin.H{"status": "1"})
}

func checkPassword(u *UserController, user models.User, password []byte) bool {
	stats := util.ComparePasswords(user.Password, password)
	return stats
}

func setPassword(u *UserController, user models.User, password []byte) {
	user.Password = util.HashAndSalt(password)
	u.userQuery.UpdateUser(user)
}
