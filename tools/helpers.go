package tools

import (
	"encoding/json"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err.Error())
}
