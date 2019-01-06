package main

import (
	"common"
	"config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var err error

	config.CONFIG, err = common.ReadConfing()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	initRoutes()
	fmt.Println("waiting routes...")
	log.Panic(http.ListenAndServe(config.CONFIG.Service.Port, app.Router))
}
