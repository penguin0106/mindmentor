package main

import (
	"github.com/gorilla/mux"
	"log"
	"mindmentor/auth_service/handlers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
