package helper 

import (
	"encoding/json"
	"net/http"
)

func WriteUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_=json.NewEncoder(w).Encode(map[string]interface{}{
		"code" : http.StatusUnauthorized,
		"status": "Unauthorized",
		"error": message,
	})
}