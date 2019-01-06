package main

import (
	"controllers"
	"structs"

	"github.com/gorilla/mux"
)

var app structs.App

func initRoutes() {
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/heartbeat", controllers.Heartbeat).Methods("GET")
	app.Router.HandleFunc("/", controllers.GetPaged).Methods("GET")
	app.Router.HandleFunc("/{id}", controllers.Get).Methods("GET")
	app.Router.HandleFunc("/{id}", controllers.Put).Methods("PUT")
	app.Router.HandleFunc("/{id}", controllers.Delete).Methods("DELETE")
}
