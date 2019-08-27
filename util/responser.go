package util

import (
	"encoding/json"
	"net/http"
)

func SendMsg(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}
