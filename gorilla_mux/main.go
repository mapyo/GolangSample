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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{name}", userHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", r))
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
