package controllers

import (
	"common"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	bookService "services/book"
	"strconv"
	"structs"

	"github.com/gorilla/mux"
)

// GetPaged is responsible for handling / GET HTTP request
func GetPaged(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	header := r.Header.Get("X-CurrentPage")
	currentPage, err := strconv.Atoi(header)

	if err != nil {
		err = errors.New("X-CurrentPage param is required")
		common.CreateBadRequestResponse(w, err)
		return
	}

	header = r.Header.Get("X-PageSize")
	pageSize, err := strconv.Atoi(header)

	if err != nil {
		err = errors.New("X-PageSize param is required")
		common.CreateBadRequestResponse(w, err)
		return
	}

	result, err := bookService.GetPaged(pageSize, currentPage)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	if len(result) == 0 {
		common.CreateEmptyResponse(w)
	}

	common.CreateSuccessResponse(w, result)
}

// Get is responsible for handling /{id} GET HTTP request
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	result, err := bookService.GetById(id)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	if result == (structs.Book{}) {
		common.CreateEmptyResponse(w)
	}

	common.CreateSuccessResponse(w, result)

}

// Post is responsible for handling / POST HTTP request
func Post(w http.ResponseWriter, r *http.Request) {
	var err error
	var book structs.Book

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	err = json.Unmarshal(body, &book)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	result, err := bookService.Save(book)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	common.CreateCreatedResponse(w, result)
}

// Put is responsible for handling /{id} PUT HTTP request
func Put(w http.ResponseWriter, r *http.Request) {
	var book structs.Book
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		err = errors.New("id param is required and must be greater than 0")
		common.CreateBadRequestResponse(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &book)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	book.Id = id
	_, err = bookService.Save(book)

	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	common.CreateSuccessResponse(w, nil)

}

// Delete is responsible for handling /{id} DELETE HTTP request
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["id"]

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		err = errors.New("id param is required and must be greater than 0")
		common.CreateBadRequestResponse(w, err)
		return
	}

	err = bookService.Remove(id)
	if err != nil {
		common.CreateBadRequestResponse(w, err)
		return
	}

	common.CreateSuccessResponse(w, nil)
}
