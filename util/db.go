package util

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *sql.DB

// GetDB method returns a DB instance
func GetDB() (*gorm.DB, error) {
	dat, err := ioutil.ReadFile("../dbconf.txt")
	if err != nil {
		panic(fmt.Sprintf("Read conf: %v", err))
	}

	conn, err := gorm.Open("postgres", string(dat))
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	return conn, nil
}
