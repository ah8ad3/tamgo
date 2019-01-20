package auth

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	templateDir = "assets/template/auth/"
)


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

func Logout(w http.ResponseWriter, r *http.Request) {
	noCache(w)
	doLogout(w, r)
	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

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
