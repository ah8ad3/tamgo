package common

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func SetDB(db *gorm.DB) {
	DB = db
}