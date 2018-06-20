package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"web/handler"
)


func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.PingHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{name}", handler.UserHandler).Methods(http.MethodGet)
	r.HandleFunc("/", handler.BasicHandler).Methods(http.MethodGet)

	return r
}
