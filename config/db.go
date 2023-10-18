package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pos-app/models"
)

var database *gorm.DB
var e error

func DatabaseInit () {

	dsn := "root:blank111@tcp(127.0.0.1:3306)/echo_rest?charset=utf8mb4&parseTime=True&loc=Local"
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	InitialMigration()
}

func DB() *gorm.DB {
	return database
}

func InitialMigration() {
	database.AutoMigrate(&models.Book{})
}

