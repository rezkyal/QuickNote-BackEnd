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

func (u *UserQuery) Init(db *gorm.DB) {
	u.db = db
}

func (u *UserQuery) CreateUser(username string) (models.User, bool) {
	var user models.User
	state := u.db.Where("Username = ?", username).First(&user)
	if gorm.IsRecordNotFoundError(state.Error) {
		user = models.User{Username: username, Password: "", CreatedOn: time.Now()}
		err := u.db.Create(&user)

		if err.Error != nil {
			log.Panic(err.Error)
		}

		return user, true
	}
	return models.User{}, false
}

func (u *UserQuery) FindOrCreateUser(username string) models.User {
	var user models.User
	state := u.db.Where("Username = ?", username).Preload("NotesOwned").First(&user)
	if state.Error != nil {
		if gorm.IsRecordNotFoundError(state.Error) {
			user = models.User{Username: username, Password: "", CreatedOn: time.Now()}
			err := u.db.Create(&user)

			if err.Error != nil {
				log.Panic(err.Error)
			}

		} else {
			log.Panic(state.Error)
		}
	}
	return user
}

func (u *UserQuery) UpdateUser(user models.User) models.User {
	err := u.db.Save(&user)
	if err.Error != nil {
		log.Panic(err.Error)
	}
	return user
}
