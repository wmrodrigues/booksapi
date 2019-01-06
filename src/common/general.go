package common

import (
	"config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
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
