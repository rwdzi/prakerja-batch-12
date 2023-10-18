package models

import "gorm.io/gorm"

type Book struct {
	ID		uint	`gorm:"primaryKey" json:"id"`
	Name 	string	`json:"name"`
	Address string 	`json:"address"`
	Phone	string  `json:"phone"`
	gorm.Model
}
