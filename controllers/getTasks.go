package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MaidedPonce/TODO-LIST-GO/repository"
	"github.com/MaidedPonce/TODO-LIST-GO/server"
)

func ReadTasks(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := repository.ReadTasks(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}
