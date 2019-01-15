package settings

import (
	"github.com/ah8ad3/tamgo/settings/db"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
)

func Config() (DB *gorm.DB){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open db and pass it to server for share it
	DB = db.OpenDB()
	return DB
}
