package common

import (
	"config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ReadConfing reads the config.json file and load the settings on memory
func ReadConfing() (config.Configuration, error) {
	var config config.Configuration

	jsonFile, err := os.Open("./config/config.json")
	defer jsonFile.Close()

	if err != nil {
		err = errors.New("error opening config.json file, " + err.Error())
		log.Println(err)
		return config, err
	}

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		err = errors.New("error reading config.json file data, " + err.Error())
		log.Println(err)
		return config, err
	}

	json.Unmarshal(bytes, &config)
	return config, nil
}

// PanicIfNotNil just panic the error on stack
func PanicIfNotNil(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// CreateBadRequestResponse creates an HTTP response with 400 HTTP status code
func CreateBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	bytes, _ := json.Marshal(err.Error())
	w.Write(bytes)
}

// CreateSuccessResponse creates an HTTP response with 200 HTTP status code
func CreateSuccessResponse(w http.ResponseWriter, data interface{}) {
	if data != nil {
		bytes, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Write(bytes)
	}
}

// CreateCreatedResponse creates an HTTP response with 201 HTTP status code
func CreateCreatedResponse(w http.ResponseWriter, data interface{}) {
	if data != nil {
		bytes, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(bytes)
	}
}

// CreateEmptyResponse creates an HTTP response with 201 HTTP status code
func CreateEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
