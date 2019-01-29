package settings

import (
	"github.com/ah8ad3/tamgo/lib/auth"
	"github.com/ah8ad3/tamgo/settings/db"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// this is config function for tamgo
func Config() (DB *gorm.DB){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store := sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options = &sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 2, // 2 hour
		HttpOnly: true,
	}

	// set session store for request
	auth.SetStore(store)

	// open db and pass it to server for share it
	DB = db.OpenDB()
	return DB
}
