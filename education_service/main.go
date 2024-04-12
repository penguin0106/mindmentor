package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mindmentor/education_service/handlers"
)

func main() {
	router := mux.NewRouter()

	// Установка обработчика для эндпоинта получения информации об образовании
	router.HandleFunc("/education", handlers.GetEducation).Methods("GET")

	log.Fatal(http.ListenAndServe(":8086", router))
}
