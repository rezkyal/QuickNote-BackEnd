package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *sql.DB

func dbConn(fileDirection string) (*gorm.DB, error) {

	conn, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	return conn, nil
}

// GetDB method returns a DB instance
func GetDB() (*gorm.DB, error) {
	return dbConn("dbconf.txt")
}

func GetDBTest() (*gorm.DB, error) {
	return dbConn("./../dbconf.txt")
}
