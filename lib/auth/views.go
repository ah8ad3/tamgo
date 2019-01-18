package auth

import (
	"fmt"
	"github.com/ah8ad3/tamgo/apps/common"
	"net/http"
)


func Register(w http.ResponseWriter, r *http.Request) {
	if res := checkLogin(r); res {
		http.Redirect(w, r, "/auth/profile", http.StatusOK)
	} else {
		_, _ = fmt.Fprintf(w, "register lodaed")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// all login staff check
	doLogin(w, r)
	_ = common.DB
	_, _ = fmt.Fprintf(w, "ok set")
}

