package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// http.HandleFunc("/", homeHandler)
	router.HandleFunc("/", homeHandler).Methods("GET")
	// escuchar y servir
	http.ListenAndServe(":3000", router)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Holu"))
}
