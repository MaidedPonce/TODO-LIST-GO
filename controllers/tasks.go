package controllers

import (
	"encoding/json"
	"net/http"

	tasks "github.com/MaidedPonce/TODO-LIST-GO/models"
	"github.com/MaidedPonce/TODO-LIST-GO/repository"
	"github.com/MaidedPonce/TODO-LIST-GO/server"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type TaskCreateReq struct {
	Text string `json:"text"`
}

type TaskCreateRes struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type TaskUpdateRes struct {
	Message string `json:"text"`
}

func CreateTask(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = TaskCreateReq{}
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
		json.NewEncoder(w).Encode(TaskCreateRes{
			Id:   task.Id,
			Text: task.Text,
		})
	}
}

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

func ReadTaskID(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskID, err := repository.ReadTaskID(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(taskID)
	}
}

func UpdateTask(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var req = TaskCreateReq{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var task = tasks.Task{
			Id:   params["id"],
			Text: req.Text,
		}

		err = repository.UpdateTask(r.Context(), &task)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TaskUpdateRes{
			Message: "se actualizo la cosa",
		})
	}
}

func DeleteTask(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := repository.DeleteTask(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TaskUpdateRes{
			Message: "se borro la cosa",
		})
	}
}
