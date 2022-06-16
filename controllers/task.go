package controllers

import (
	"encoding/json"
	"net/http"

	tasks "github.com/MaidedPonce/TODO-LIST-GO/models"
	"github.com/MaidedPonce/TODO-LIST-GO/repository"
	"github.com/MaidedPonce/TODO-LIST-GO/server"
	"github.com/segmentio/ksuid"
)

type TaskReq struct {
	Text string `json:"text"`
}

type TaskRes struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func CreateTask(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = TaskReq{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var task = tasks.Task{
			Id:   id.String(),
			Text: req.Text,
		}

		err = repository.CreateTask(r.Context(), &task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TaskRes{
			Id:   task.Id,
			Text: task.Text,
		})
	}
}
