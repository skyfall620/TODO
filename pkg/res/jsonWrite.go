package res

import (
	"encoding/json"
	"net/http"
)

func JSONWrite(w http.ResponseWriter, data any, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}
