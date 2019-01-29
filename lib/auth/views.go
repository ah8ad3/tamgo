package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

const (
	templateDir = "assets/template/auth/"
)


// Register user view
func Register(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	if r.Method == "GET" {
		res := checkLogin(r)
		if res == true {
			http.Redirect(w, r, "/auth/profile", http.StatusFound)
			return
		} else if res == false {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			http.ServeFile(w, r, templateDir + "register.html")
			return
		}
	}else {
		_ = r.ParseForm()
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")
		email := strings.Join(r.Form["email"], "")

		if username == "" || password == "" || email == "" {
			_, _ = fmt.Fprintf(w, "input form are incomplete")
			return
		}

		user := User{}
		DB.Where("username = ?", username).First(&user)
		if user.ID != 0 {
			_, _ = fmt.Fprintf(w, "this username taken")
			return
		}

		DB.Create(&User{Username: username, Password: SetPassword(password), Email: email})

		http.Redirect(w, r, "/auth/login", http.StatusFound)
		return
	}
}

// Login create session for login
func Login(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, templateDir + "login.html")
		return
	}else {
		_ = r.ParseForm()
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")

		if username == "" || password == "" {
			_, _ = fmt.Fprintf(w, "input form are incomplete")
			return
		}

		user := User{}
		DB.Where("username = ?", username).First(&user)

		if user.ID == 0 || !ValidPassword(user.Password, password){
			_, _ = fmt.Fprintf(w, "username or password wrong")
			return
		}

		doLogin(w, r)

		http.Redirect(w, r, "/auth/profile", http.StatusFound)

		return
	}
}

// Logout function for session
func Logout(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	doLogout(w, r)
	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

// Profile view for user profile an secret area
func Profile(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	res := checkLogin(r)
	if res == true {
		http.ServeFile(w, r, templateDir+"profile.html")
		return
	}else {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	}
}

// CheckJwt if token is JWT and have time
func CheckJwt(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	tokenstring := strings.Join(r.Form["token"], "")
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	// When using `Parse`, the result `Claims` would be a map.

	// In another way, you can decode token to your struct, which needs to satisfy `jwt.StandardClaims`
	user := JWT{}
	token, _ = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if token.Valid {
		_, _ = fmt.Fprintf(w, "this token is right")
		return
	}
	_, _ = fmt.Fprintf(w, "this token is wrong")
	return
}

// SignJWT function to create jwt token
func SignJWT(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")

		if username == "" || password == "" {
			_, _ = fmt.Fprintf(w, "input form are incomplete")
			return
		}

		user := User{}
		DB.Where("username = ?", username).First(&user)

		if user.ID == 0 || !ValidPassword(user.Password, password){
			_, _ = fmt.Fprintf(w, "username or password wrong")
			return
		}
		token := createTokenString(user)
		_, _ = fmt.Fprintf(w, token)
		return
	}
}
