package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPaged is responsible for handling / GET HTTP request
func GetPaged(w http.ResponseWriter, r *http.Request) {

}

// Get is responsible for handling /{id} GET HTTP request
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(id)
}

// Post is responsible for handling / POST HTTP request
func Post(w http.ResponseWriter, r *http.Request) {

}

// Put is responsible for handling /{id} PUT HTTP request
func Put(w http.ResponseWriter, r *http.Request) {

}

// Delete is responsible for handling /{id} DELETE HTTP request
func Delete(w http.ResponseWriter, r *http.Request) {

}
