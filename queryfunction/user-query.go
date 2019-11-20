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

func (n *UserQuery) FindOrCreateUser(username string) models.User {
	var user models.User
	state := n.db.Where("Username = ?", username).First(&user)
	if state.Error != nil {
		if gorm.IsRecordNotFoundError(state.Error) {
			user = models.User{Username: username, Password: "", CreatedOn: time.Now()}
			err := n.db.Create(&user)

			if err.Error != nil {
				log.Panic(err.Error)
			}

		} else {
			log.Panic(state.Error)
		}
	}
	return user
}
