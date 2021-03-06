package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/MaidedPonce/TODO-LIST-GO/controllers"
	"github.com/MaidedPonce/TODO-LIST-GO/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("hubo un error en env:(")
	}

	PORT := os.Getenv("PORT")
	DATABASE := os.Getenv("DATABASE")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		DataBaseUrl: DATABASE,
	})

	if err != nil {
		log.Fatal(err)
	}
	s.Start(BindRputes)
}

func BindRputes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", controllers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/newtask", controllers.CreateTask(s)).Methods(http.MethodPost)
	r.HandleFunc("/tasks", controllers.ReadTasks(s)).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", controllers.ReadTaskID((s))).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask(s)).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask(s)).Methods(http.MethodDelete)
}
