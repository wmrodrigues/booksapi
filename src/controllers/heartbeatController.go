package controllers

import (
	"encoding/json"
	"net/http"
	"structs"
)

// Heartbeat is responsible for hanlding the /heartbeat GET HTTP request
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := structs.HeartbeatResponse{Code: 200, Status: "OK"}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
