package auth

import (
	"fmt"
	"github.com/ah8ad3/tamgo/apps/common"
	"net/http"
)


func MyHandler(w http.ResponseWriter, r *http.Request) {
	session1, _ := Store.Get(r, "loginSess1")
	session1.Values["foo"] = "bar"
	_ = session1.Save(r, w)
	_ = common.DB
	_, _ = fmt.Fprintf(w, "ok set")
}


func GetFoo(f interface{}) string {
	if f != nil {
		if foo, ok := f.(string); ok {
			return foo
		}
	}
	return ""
}


func SessionCheck(w http.ResponseWriter, r *http.Request) {
	session1, _ := Store.Get(r, "loginSess1")

	if session1.Values["foo"] == nil {
		fmt.Println("session not exist")
	}

	_, _ = fmt.Fprintf(w, GetFoo(session1.Values["foo"]))
}
