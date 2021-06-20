package main

import (
	"github.com/gorilla/mux"
	"github.com/yubing24/golang-angular-fullstack-2021/api"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Methods(http.MethodGet, http.MethodOptions).Path("/search-country").HandlerFunc(api.ResponseInterceptor(api.SearchCountry))
	http.Handle("/", router)
	log.Fatalf("[fatal] %s", http.ListenAndServe(":8080", nil))
}