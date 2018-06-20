package handler

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

var DB = map[string]string{
	"user1": "foo",
	"user2": "bar",
}

var secrets = map[string]string{
	"foo":  "secret1",
	"manu": "seecret2",
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["name"]

	value, ok := DB[user]
	if ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, value)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || !existsSecretUser(user, pass) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "user: %s, secret: %s", user, pass)
}

func existsSecretUser(user string, pass string) bool {
	if _, ok := secrets[user]; ok {
		return true
	} else {
		return false
	}
}
