package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit () {
	hot := "localhost"
	user := "root"
	password := "blank111"
	dbName : "echo_rest"
	port := 1212

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panicc(e)
	}
}

func DB() *gorm.DB {
	return database
}

