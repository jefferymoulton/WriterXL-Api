package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "writerxl:admin123@tcp(localhost:3306)/writerxl?charset=utf8&parseTime=True")

	if err != nil {
		errMsg := "Failed to connect to the database "
		panic(errMsg + err.Error())
	}

	database.AutoMigrate(&User{})

	database.LogMode(true)

	DB = database
}
