package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

const (
	// yellow color for warning
	WarningColor = "\033[1;33m%s\033[0m"
)

// connect to pq database with gorm
func OpenDB() (db *gorm.DB) {
	fmt.Println(os.Getenv("DB_HOST"))

	DB, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS")))
	//DB, err := gorm.Open("postgres", fmt.Sprintf(
	//	"postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	_ = DB

	fmt.Printf(WarningColor, "database connected")
	fmt.Println("")

	return DB
}
