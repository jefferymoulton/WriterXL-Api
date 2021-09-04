package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {

	database, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION"))

	if err != nil {
		errMsg := "Failed to connect to the database "
		panic(errMsg + err.Error())
	}

	database.AutoMigrate(&Profile{})

	database.LogMode(true)

	DB = database
}
