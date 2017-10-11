package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest-in-go/api"
)

var routes = mux.NewRouter().PathPrefix("/api/v1").Subrouter()

// Route Configuration
func initRoutes() {
	routes.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	http.ListenAndServe(":8009", routes)
}

// Application Start Point
func main() {
	initRoutes()
}
