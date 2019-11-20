package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/rezkyal/QuickNote-BackEnd/queryfunction"
)

type UserController struct {
	userQuery *queryfunction.UserQuery
}

func (u *UserController) Init(db *gorm.DB) {
	u.userQuery = &queryfunction.UserQuery{}
	u.userQuery.Init(db)
}
