package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

const (
	WarningColor = "\033[1;33m%s\033[0m"
)

func OpenDB() (db *gorm.DB) {

	DB, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS")))

	// defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = DB

	fmt.Printf(WarningColor, "database connected")
	fmt.Println("")

	return DB
}