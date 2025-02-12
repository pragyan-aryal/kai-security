package v1

import (
	"encoding/json"
	"net/http"

	"kai_security/internal/service/scan"
)

func ScanGithub(w http.ResponseWriter, r *http.Request) {
	var scanRequest scan.Request

	err := json.NewDecoder(r.Body).Decode(&scanRequest)
	if err != nil {
		JSONErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}

func QueryVulerabilities(w http.ResponseWriter, r *http.Request) {
	var scanRequest scan.Request

	err := json.NewDecoder(r.Body).Decode(&scanRequest)
	if err != nil {
		JSONErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}

func JSONErrorResponse(rw http.ResponseWriter, err error, statusCode int) {
	rw.WriteHeader(statusCode)
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{"error": err})
}

func JSONSuccessResponse(rw http.ResponseWriter, data interface{}, statusCode int) {
	rw.WriteHeader(statusCode)
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(data)
}
