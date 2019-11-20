package queryfunction

import (
	"log"
	"time"

	"github.com/rezkyal/QuickNote-BackEnd/models"

	"github.com/jinzhu/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func (n *UserQuery) Init(db *gorm.DB) {
	n.db = db
}

func (n *UserQuery) FindOrCreateUser(username string) int64 {
	var user models.User
	state := n.db.Where("Username = ?", username).First(&user)
	if state.Error != nil {
		if gorm.IsRecordNotFoundError(state.Error) {
			user = models.User{Username: username, Password: "", CreatedOn: time.Now()}
			n.db.Create(&user)
		} else {
			log.Panic(state.Error)
		}
	}

	return user.UserID
}
