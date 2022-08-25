package controller

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
}
