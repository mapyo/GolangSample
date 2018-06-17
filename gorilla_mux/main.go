package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8081", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "mux sample")
}
