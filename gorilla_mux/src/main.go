package main

import (
	"net/http"
	"log"
	"web"
)

func main() {
	log.Fatal(http.ListenAndServe(":8082", web.Router()))
}



