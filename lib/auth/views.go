package auth

import (
	"fmt"
	"net/http"
)

const (
	templateDir = "assets/template/auth/"
)


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if res := checkLogin(r); res {
			http.Redirect(w, r, "/auth/profile", http.StatusFound)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			http.ServeFile(w, r, templateDir + "register.html")
		}
	}else {
		_, _ = fmt.Fprintf(w, "post of register section")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, templateDir + "login.html")
	}else {
		// all login staff check
		//doLogin(w, r)
		_, _ = fmt.Fprintf(w, "post section of login")
	}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	isLoggedIn(w, r)
	_, _ = fmt.Fprintf(w, "this is secret area")
}

func isLoggedIn(w http.ResponseWriter, r *http.Request) {
	if res := checkLogin(r); res == false {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	}
}
