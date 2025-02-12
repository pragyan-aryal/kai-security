package v1

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/resource", ScanGithub).Methods("POST")
	v1.HandleFunc("/resource", QueryVulerabilities).Methods("POST")
}
