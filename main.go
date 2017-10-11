package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/s4kibs4mi/rest-in-go/api"
)

var routes = mux.NewRouter()

// Route Configuration
func initRoutes() {
	// Parent Routes
	routes.HandleFunc("/", api.IndexHandler).Methods("GET", "POST", "PATCH", "OPTIONS", "DELETE", "PUT")

	// Sub-routes { /api/v1 }
	v1 := routes.NewRoute().PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", api.CreateUser).Methods("POST")
	v1.HandleFunc("/user/update", api.Authorization(api.UpdateUser)).Methods("PUT") // Unimplemented
	v1.HandleFunc("/user/auth/login", api.LoginUser).Methods("POST")
	v1.HandleFunc("/user/auth/logout/{access_token}", api.DeleteAccessToken).Methods("DELETE") // Unimplemented

	v1.HandleFunc("/contact/create", api.Authorization(api.CreateContact)).Methods("POST")
	v1.HandleFunc("/contact/list", api.Authorization(api.ListContact)).Methods("GET")
	v1.HandleFunc("/contact/delete/{contact_id}", api.Authorization(api.DeleteContact)).Methods("DELETE") // Unimplemented
	v1.HandleFunc("/contact/update/{contact_id}", api.Authorization(api.UpdateContact)).Methods("PUT")    // Unimplemented

	fmt.Println("Application Running On :8009...")
	http.ListenAndServe(":8009", routes)
}

// Application Start Point
func main() {
	initRoutes()
}
