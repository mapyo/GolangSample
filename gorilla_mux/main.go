package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

var DB = map[string]string{
	"user1": "foo",
	"user2": "bar",
}

var secrets = map[string]string{
	"foo":  "secret1",
	"manu": "seecret2",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{name}", userHandler).Methods(http.MethodGet)
	r.HandleFunc("/", basicHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8082", r))
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
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

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["name"]

	value, ok := DB[user]
	if ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, value)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}
