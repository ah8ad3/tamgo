package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)


func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		//log.Println(err)
		return false
	}

	return true
}

func getAuthInterface(f interface{}) bool {
	if f != nil {
		if ok, _ := f.(bool); ok {
			return true
		}
	}
	return false
}

func checkLogin(r *http.Request) bool {
	ses, _ := Store.Get(r, "tamgo-auth")
	if res := getAuthInterface(ses.Values["login"]); res {
		return true
	}
	return false
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	ses, _ := Store.Get(r, "tamgo-auth")
	ses.Values["login"] = true
	_ = ses.Save(r, w)
}

func doLogout(w http.ResponseWriter, r *http.Request) {
	ses, _ := Store.Get(r, "tamgo-auth")
	ses.Values["login"] = false
	_ = ses.Save(r, w)
}

func noCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
}
