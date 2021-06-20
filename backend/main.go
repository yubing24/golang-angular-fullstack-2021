package main

import (
	"github.com/yubing24/golang-angular-fullstack-2021/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.SearchCountry)
	log.Fatalf("[fatal] %s", http.ListenAndServe(":8080", nil))
}