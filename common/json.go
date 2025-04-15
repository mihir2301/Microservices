package common

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func Readjson(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(&data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJson(w, status, map[string]string{"error": message})
}
