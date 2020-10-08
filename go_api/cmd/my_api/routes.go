package main

import (
	"encoding/json"
	"net/http"
	// "github.com/drewburns/my_api/internal/schema" //This loads our schema
)

func Sanity(w http.ResponseWriter, r *http.Request) {
	// tokenVals := r.Context().Value("user") // Get JWT
	// r.URL.Query().Get("private") // Get query
	// params := mux.Vars(r) // Get query

	// Error return:
	// w.WriteHeader(http.StatusBadRequest)
	// w.Write([]byte("Public posts can only be less than 50 characters."))

	json.NewEncoder(w).Encode("OK")
}
